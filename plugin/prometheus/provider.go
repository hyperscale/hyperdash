package main

import (
	"github.com/hyperscale/hyperdash/pkg/hyperdash/provider"
	"github.com/hyperscale/hyperdash/plugin/prometheus/prometheus"
)

// Provider return provider instance
func Provider() provider.Provider {
	return prometheus.New()
}
