package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/euskadi31/go-sse"
	"github.com/google/uuid"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/bus"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/config"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/protocol"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/provider"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/ui"
	"github.com/hyperscale/hyperdash/plugin/apijson/apijson"
	"github.com/hyperscale/hyperdash/plugin/healthcheck/healthcheck"
	"github.com/hyperscale/hyperdash/plugin/prometheus/prometheus"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// RunAction handle run http server
func RunAction(c *cli.Context) error {
	var previousMetrics sync.Map

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	parser := config.NewParser(nil)

	file, diags := parser.LoadConfigFile(c.Args().First())
	if diags.HasErrors() {
		log.Fatal().Err(diags).Msg("")
	}

	file.Version = time.Now().Unix()

	spew.Dump(file)

	providerMap := make(map[string]provider.Provider)

	messages := bus.New()

	go messages.Start()

	pmch := make(chan *protocol.Message, 10)

	defer close(pmch)

	if err := messages.Register("previous-metric", pmch); err != nil {
		return err
	}

	defer messages.Unregister("previous-metric")

	go func(pmch chan *protocol.Message) {
		for t := range pmch {
			if t.Target != "" {
				previousMetrics.Store(t.Target, t)
			}
		}
	}(pmch)

	providers := provider.NewContainer()

	if err := providers.AddProvider(func() provider.Provider {
		return healthcheck.New()
	}); err != nil {
		return err
	}

	if err := providers.AddProvider(func() provider.Provider {
		return prometheus.New()
	}); err != nil {
		return err
	}

	if err := providers.AddProvider(func() provider.Provider {
		return apijson.New()
	}); err != nil {
		return err
	}

	if file.ProviderDir != "" {
		if err := providers.LoadProviderDir(file.ProviderDir); err != nil {
			return err
		}
	}

	// run provider
	for _, tile := range file.Tiles {
		p, err := providers.Get(tile.Provider.Name)
		if err != nil {
			log.Error().Err(err).Msgf("")

			continue
		}

		if _, ok := providerMap[tile.Name]; ok {
			log.Error().Err(fmt.Errorf("cannot override instance of provider for %s name", tile.Name)).Msg("")

			continue
		}

		providerMap[tile.Name] = p

		go func(tile *config.Tile, p provider.Provider) {
			if err := p.Start(tile, messages); err != nil {
				log.Error().Err(err).Msg("")
			}
		}(tile, p)
	}

	http.Handle("/scripts/", ui.Scripts())
	http.Handle("/styles/", ui.Styles())

	serve := sse.NewServer(func(rw sse.ResponseWriter, r *http.Request) {
		ch := make(chan *protocol.Message, 10)
		id := uuid.New().String()

		if err := messages.Register(id, ch); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)

			return
		}

		defer func() {
			messages.Unregister(id)

			close(ch)
		}()

		v := r.URL.Query().Get("v")

		version, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Error().Err(err).Msg("failed to parse version from query string")
		}

		updateNotified := false

		if version < file.Version && updateNotified == false {
			ch <- &protocol.Message{
				Type: protocol.MessageTypeUpdate,
			}

			updateNotified = true
		}

		log.Debug().Msg("sending previous data")
		previousMetrics.Range(func(key interface{}, value interface{}) bool {
			val := value.(*protocol.Message)

			b, err := json.Marshal(val)
			if err != nil {
				log.Error().Err(err).Msg("marshaling message failed")
				return true
			}

			log.Debug().Msgf("sent %s metric", val.Target)
			rw.Send(&sse.MessageEvent{
				Data: b,
			})

			return true
		})

		for {
			select {
			case t := <-ch:
				log.Debug().Msgf("Send event %s to %s...", t.Type, t.Target)

				b, err := json.Marshal(t)
				if err != nil {
					log.Error().Err(err).Msg("marshaling message failed")
					continue
				}

				rw.Send(&sse.MessageEvent{
					Data: b,
				})
			case <-r.Context().Done():
				log.Debug().Msg("Done")

				return
			}
		}
	})

	serve.SetRetry(time.Second * 5)

	http.Handle("/events", serve)

	index := ui.NewEngine(file)

	http.Handle("/", index)

	go func() {
		if err := http.ListenAndServe(":"+strconv.Itoa(file.Dashboard.Server.Port), nil); err != nil {
			log.Error().Err(err).Msg("")
		}
	}()

	<-sig

	for n, p := range providerMap {
		if err := p.Stop(); err != nil {
			log.Error().Err(err).Msgf("stop provider %q failed", n)
		}
	}

	messages.Stop()

	return nil
}
