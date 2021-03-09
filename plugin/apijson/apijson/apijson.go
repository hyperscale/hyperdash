package apijson

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/bus"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/config"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/protocol"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/provider"
	"github.com/rs/zerolog/log"
	"github.com/savaki/jq"
)

var _ provider.Provider = (*Provider)(nil)

// Provider struct
type Provider struct {
	ticker *time.Ticker
	done   chan bool
}

// New provider
func New() provider.Provider {
	return &Provider{
		done: make(chan bool),
	}
}

// Info return provider info
func (p *Provider) Info() provider.Info {
	return provider.Info{
		Name:        "apijson",
		Version:     "1.0.0",
		Description: "API JSON provider",
		Author:      "Axel Etcheverry",
	}
}

func (p *Provider) fetch(cfg *Config, client *http.Client) ([]byte, error) {
	log.Debug().Str("url", cfg.URL).Msg("calling...")
	req, err := http.NewRequest(http.MethodGet, cfg.URL, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range cfg.Headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (p *Provider) process(cfg *Config, httpClient *http.Client, tile *config.Tile, messages bus.Writer) {
	body, err := p.fetch(cfg, httpClient)
	if err != nil {
		log.Error().Err(err).Msg("")

		messages.Publish(&protocol.Message{
			Target: tile.Name,
			Type:   protocol.MessageTypeError,
			Payload: &protocol.MessageError{
				Level: protocol.ErrorLevelTypeError,
			},
		})

		return
	}

	op, err := jq.Parse(cfg.Query)
	if err != nil {
		log.Error().Err(err).Msg("")

		messages.Publish(&protocol.Message{
			Target: tile.Name,
			Type:   protocol.MessageTypeError,
			Payload: &protocol.MessageError{
				Level: protocol.ErrorLevelTypeError,
			},
		})

		return
	}

	value, err := op.Apply(body)
	if err != nil {
		log.Error().Err(err).Msg("")

		messages.Publish(&protocol.Message{
			Target: tile.Name,
			Type:   protocol.MessageTypeError,
			Payload: &protocol.MessageError{
				Level: protocol.ErrorLevelTypeError,
			},
		})

		return
	}

	val, err := strconv.ParseFloat(string(value), 64)
	if err != nil {
		log.Error().Err(err).Msg("")

		messages.Publish(&protocol.Message{
			Target: tile.Name,
			Type:   protocol.MessageTypeError,
			Payload: &protocol.MessageError{
				Level: protocol.ErrorLevelTypeError,
			},
		})

		return
	}

	messages.Publish(&protocol.Message{
		Target: tile.Name,
		Type:   protocol.MessageTypeStat,
		Payload: &protocol.MessageStat{
			Value: val,
			Unit:  tile.Unit.String(),
		},
	})
}

// Start func
func (p *Provider) Start(tile *config.Tile, messages bus.Writer) error {
	cfg := &Config{}

	if diags := gohcl.DecodeBody(tile.Provider.Body, nil, cfg); diags.HasErrors() {
		return diags
	}

	var httpClient = &http.Client{
		Timeout: cfg.TimeoutDuration(),
	}

	p.ticker = time.NewTicker(cfg.IntervalDuration())

	log.Debug().
		Dur("interval", cfg.IntervalDuration()).
		Dur("timeout", cfg.TimeoutDuration()).
		Msgf("Start %s %s", p.Info().Name, p.Info().Version)

	time.Sleep(5 * time.Second)

	p.process(cfg, httpClient, tile, messages)

	for {
		select {
		case <-p.done:
			return nil
		case <-p.ticker.C:
			p.process(cfg, httpClient, tile, messages)
		}
	}
}

// Stop func
func (p *Provider) Stop() error {

	p.ticker.Stop()

	p.done <- true

	return nil
}
