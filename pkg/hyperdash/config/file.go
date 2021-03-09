package config

// File struct
type File struct {
	Version     int64
	ProviderDir string     `hcl:"provider_dir,optional"`
	Dashboard   *Dashboard `hcl:"dashboard,block"`
	Tiles       []*Tile    `hcl:"tile,block"`
}
