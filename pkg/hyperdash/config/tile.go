package config

// Tile struct
type Tile struct {
	Type     string    `hcl:"type,label"`
	Name     string    `hcl:"name,label"`
	Column   int       `hcl:"column"`
	Row      int       `hcl:"row"`
	Title    string    `hcl:"title"`
	Unit     UnitType  `hcl:"unit,optional"`
	Provider *Provider `hcl:"provider,block"`
}
