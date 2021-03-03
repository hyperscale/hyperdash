package config

// Grid struct
type Grid struct {
	Columns int `hcl:"columns"`
	Rows    int `hcl:"rows"`
}
