package config

// Tile struct
type Tile struct {
	Type     string    `hcl:"type,label"`
	Name     string    `hcl:"name,label"`
	Column   int       `hcl:"column"`
	Row      int       `hcl:"row"`
	Provider *Provider `hcl:"provider,block"`
}
