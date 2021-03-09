package healthcheck

import (
	"net/http"
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
		Name:        "healthcheck",
		Version:     "1.0.0",
		Description: "Healthcheck http service",
		Author:      "Axel Etcheverry",
	}
}

func (p *Provider) check(cfg *Config, client *http.Client) bool {
	resp, err := client.Get(cfg.URL)
	if err != nil {
		return false
	}

	return resp.StatusCode == 200
}

func (p *Provider) getStatus(buffer *CircularStatus) protocol.StatusType {
	sum := 0

	for _, v := range buffer.Values() {
		sum += v
	}

	if sum == buffer.Cap() {
		return protocol.StatusTypeGreen
	} else if sum == 0 {
		return protocol.StatusTypeRed
	}

	return protocol.StatusTypeYellow
}

func (p *Provider) process(cfg *Config, httpClient *http.Client, tile *config.Tile, buffer *CircularStatus, messages bus.Writer) {
	status := 0

	if ok := p.check(cfg, httpClient); ok {
		status = 1
	}

	buffer.Push(status)

	messages.Publish(&protocol.Message{
		Target: tile.Name,
		Type:   protocol.MessageTypeStatus,
		Payload: &protocol.MessageStatus{
			Status: p.getStatus(buffer),
		},
	})
}

// Start func
func (p *Provider) Start(tile *config.Tile, messages bus.Writer) error {
	cfg := &Config{}

	if diags := gohcl.DecodeBody(tile.Provider.Body, nil, cfg); diags.HasErrors() {
		return diags
	}

	buffer := NewCircularStatus(cfg.HealthyThreshold)

	var httpClient = &http.Client{
		Timeout: cfg.TimeoutDuration(),
	}

	p.ticker = time.NewTicker(cfg.IntervalDuration())

	log.Debug().
		Dur("interval", cfg.IntervalDuration()).
		Dur("timeout", cfg.TimeoutDuration()).
		Int("unhealthy_threshold", cfg.UnhealthyThreshold).
		Int("healthy_threshold", cfg.HealthyThreshold).
		Msgf("Start %s %s", p.Info().Name, p.Info().Version)

	time.Sleep(5 * time.Second)

	p.process(cfg, httpClient, tile, buffer, messages)

	for {
		select {
		case <-p.done:
			return nil
		case <-p.ticker.C:
			p.process(cfg, httpClient, tile, buffer, messages)
		}
	}
}

// Stop func
func (p *Provider) Stop() error {

	p.ticker.Stop()

	p.done <- true

	return nil
}
