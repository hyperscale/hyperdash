package config

// Dashboard struct
type Dashboard struct {
	Title  string  `hcl:"title"`
	Grid   *Grid   `hcl:"grid,block"`
	Server *Server `hcl:"server,block"`
}
