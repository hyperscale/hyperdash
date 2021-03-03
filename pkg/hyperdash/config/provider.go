package config

import "github.com/hashicorp/hcl/v2"

// Provider struct
type Provider struct {
	Name string   `hcl:"name,label"`
	Body hcl.Body `hcl:",remain"`
}
