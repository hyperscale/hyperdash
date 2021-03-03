package provider

import "github.com/hyperscale/hyperdash/pkg/hyperdash/bus"

// Provider interface
type Provider interface {
	Info() Info
	Config() interface{}
	Start(config interface{}, messages bus.Writer, target string) error
	Stop() error
}
