package main

import (
	"github.com/hyperscale/hyperdash/pkg/hyperdash/provider"
	"github.com/hyperscale/hyperdash/plugin/apijson/apijson"
)

// Provider return provider instance
func Provider() provider.Provider {
	return apijson.New()
}
