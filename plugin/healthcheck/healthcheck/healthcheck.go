package healthcheck

import (
	"time"

	"github.com/hyperscale/hyperdash/pkg/hyperdash/bus"
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
		Name:        "healthcheck",
		Version:     "1.0.0",
		Description: "Healthcheck http service",
		Author:      "Axel Etcheverry",
	}
}

// Config return config provider
func (p *Provider) Config() interface{} {
	return &Config{}
}

// Start func
func (p *Provider) Start(config interface{}, messages bus.Writer, target string) error {
	cfg := config.(*Config)

	p.ticker = time.NewTicker(cfg.IntervalDuration())

	log.Debug().
		Dur("interval", cfg.IntervalDuration()).
		Dur("timeout", cfg.TimeoutDuration()).
		Int("unhealthy_threshold", cfg.UnhealthyThreshold).
		Int("healthy_threshold", cfg.HealthyThreshold).
		Msgf("Start %s %s", p.Info().Name, p.Info().Version)

	for {
		select {
		case <-p.done:
			return nil
		case <-p.ticker.C:
			//@TODO: exec http request

			messages.Publish(&protocol.Message{
				Target:  target,
				Type:    protocol.MessageTypeStatus,
				Payload: protocol.StatusTypeGreen,
			})
		}
	}
}

// Stop func
func (p *Provider) Stop() error {

	p.ticker.Stop()

	p.done <- true

	return nil
}
