package config

import (
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/spf13/afero"
	"github.com/zclconf/go-cty/cty"
)

// Parser struct
type Parser struct {
	fs  afero.Afero
	p   *hclparse.Parser
	ctx *hcl.EvalContext
}

func getEnvVars() map[string]cty.Value {
	envs := map[string]cty.Value{}

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)

		envs[pair[0]] = cty.StringVal(pair[1])
	}

	return envs
}

// NewParser creates and returns a new Parser that reads files from the given
// filesystem. If a nil filesystem is passed then the system's "real" filesystem
// will be used, via afero.OsFs.
func NewParser(fs afero.Fs) *Parser {
	if fs == nil {
		fs = afero.OsFs{}
	}

	envs := getEnvVars()

	return &Parser{
		fs: afero.Afero{Fs: fs},
		p:  hclparse.NewParser(),
		ctx: &hcl.EvalContext{
			Variables: map[string]cty.Value{
				"env": cty.ObjectVal(envs),
			},
		},
	}
}

// Parse config
func (p *Parser) Parse(filename string, src []byte) (*File, hcl.Diagnostics) {
	f, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{Line: 1, Column: 1})

	if diags.HasErrors() {
		return nil, diags
	}

	cfg := &File{}

	diags = gohcl.DecodeBody(f.Body, p.ctx, cfg)

	return cfg, diags
}
