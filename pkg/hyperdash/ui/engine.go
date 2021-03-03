package ui

import (
	"html/template"
	"net/http"

	"github.com/hyperscale/hyperdash/pkg/hyperdash/config"
)

// Engine struct
type Engine struct {
	tmpl *template.Template
	cfg  *config.File
}

// NewEngine constructor
func NewEngine(cfg *config.File) *Engine {
	tmpl := template.Must(template.ParseFS(index, "index.tmpl"))

	return &Engine{
		tmpl: tmpl,
		cfg:  cfg,
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := e.tmpl.Execute(w, e.cfg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
