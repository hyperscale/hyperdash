package prometheus

import (
	"encoding/json"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/bus"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/config"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/protocol"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/provider"
	"github.com/rs/zerolog/log"
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
		Name:        "prometheus",
		Version:     "1.0.0",
		Description: "Prometheus provider",
		Author:      "Axel Etcheverry",
	}
}

func (p *Provider) fetch(cfg *Config, client *http.Client) *Response {
	url := strings.TrimRight(cfg.URL, "/") + "/api/v1/query?query=" + url.QueryEscape(cfg.Query)

	resp, err := client.Get(url)
	if err != nil {
		return &Response{
			Status: "error",
		}
	}
	defer resp.Body.Close()

	data := &Response{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		log.Error().Err(err).Msg("unmarshal body failed")

		return &Response{
			Status: "error",
		}
	}

	return data
}

func (p *Provider) convertResultToMessage(tile *config.Tile, stats *Response) *protocol.Message {
	switch stats.Data.ResultType {
	case ValueTypeVector:
		var result ResultVectorType

		if err := json.Unmarshal(stats.Data.Result, &result); err != nil {
			return &protocol.Message{
				Type: protocol.MessageTypeError,
				Payload: &protocol.MessageError{
					Level: protocol.ErrorLevelTypeError,
				},
			}
		}

		metric := ResultVector{}

		if len(result) > 0 {
			metric = result[0]
		}

		var value float64
		var err error

		if len(metric.Value) == 2 {
			v := metric.Value[1].(string)

			value, err = strconv.ParseFloat(v, 64)
			if err != nil {
				log.Error().Err(err).Msg("failed to parse float value")

				return &protocol.Message{
					Type: protocol.MessageTypeError,
					Payload: &protocol.MessageError{
						Level: protocol.ErrorLevelTypeError,
					},
				}
			}

			value = math.Ceil(value*100) / 100
		}

		return &protocol.Message{
			Type: protocol.MessageTypeStat,
			Payload: &protocol.MessageStat{
				Value: value,
				Unit:  tile.Unit.String(),
			},
		}
	default:
		log.Error().Msgf("prometheus result type %q is not supported", stats.Data.ResultType)

		return &protocol.Message{
			Type: protocol.MessageTypeError,
			Payload: &protocol.MessageError{
				Level: protocol.ErrorLevelTypeError,
			},
		}
	}
}

func (p *Provider) process(cfg *Config, httpClient *http.Client, tile *config.Tile, messages bus.Writer) {
	stats := p.fetch(cfg, httpClient)

	if stats.Status != StatusTypeSuccess {
		messages.Publish(&protocol.Message{
			Target: tile.Name,
			Type:   protocol.MessageTypeError,
			Payload: &protocol.MessageError{
				Level: protocol.ErrorLevelTypeError,
			},
		})

		return
	}

	msg := p.convertResultToMessage(tile, stats)
	msg.Target = tile.Name

	messages.Publish(msg)
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
