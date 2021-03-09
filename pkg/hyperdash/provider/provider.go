package provider

import (
	"github.com/hyperscale/hyperdash/pkg/hyperdash/bus"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/config"
)

// Provider interface
type Provider interface {
	Info() Info
	Start(tile *config.Tile, messages bus.Writer) error
	Stop() error
}
