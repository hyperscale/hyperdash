package main

import (
	"github.com/hyperscale/hyperdash/pkg/hyperdash/provider"
	"github.com/hyperscale/hyperdash/plugin/healthcheck/healthcheck"
)

// Provider return provider instance
func Provider() provider.Provider {
	return healthcheck.New()
}
