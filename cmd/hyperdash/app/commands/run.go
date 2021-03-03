package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/euskadi31/go-sse"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/bus"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/config"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/provider"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/ui"
	"github.com/hyperscale/hyperdash/plugin/healthcheck/healthcheck"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// RunAction handle run http server
func RunAction(c *cli.Context) error {
	parser := config.NewParser(nil)

	file, diags := parser.LoadConfigFile(c.Args().First())
	if diags.HasErrors() {
		log.Fatal().Err(diags).Msg("")
	}

	spew.Dump(file)

	providerMap := make(map[string]provider.Provider)

	messages := bus.New()

	providers := provider.NewContainer()

	if err := providers.AddProvider(func() provider.Provider {
		return healthcheck.New()
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

		cfg := p.Config()

		if diags := gohcl.DecodeBody(tile.Provider.Body, nil, cfg); diags.HasErrors() {
			log.Error().Err(diags).Msg("")
		}

		providerMap[tile.Name] = p

		go func(cfg interface{}, p provider.Provider, target string) {
			if err := p.Start(cfg, messages, target); err != nil {
				log.Error().Err(err).Msg("")
			}
		}(cfg, p, tile.Name)
	}

	http.Handle("/scripts/", ui.Scripts())
	http.Handle("/styles/", ui.Styles())

	serve := sse.NewServer(func(rw sse.ResponseWriter, r *http.Request) {
		// recovery
		lastID := r.Header.Get(sse.LastEventID)
		if lastID != "" {
			log.Printf("Recovery with ID: %s\n", lastID)
		}

		for {
			select {
			case t := <-messages.Retrieve():
				log.Debug().Msg("Send event...")

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

	return http.ListenAndServe(":8080", nil)
}
