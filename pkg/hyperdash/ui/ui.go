package ui

import (
	"embed"
	"net/http"
)

//go:embed styles/*
var styles embed.FS

//go:embed scripts/*
var scripts embed.FS

//go:embed index.tmpl
var index embed.FS

// Styles handler
func Styles() http.Handler {
	return http.FileServer(http.FS(styles))
}

// Scripts handler
func Scripts() http.Handler {
	return http.FileServer(http.FS(scripts))
}

// Index handler
func Index() http.Handler {
	return http.FileServer(http.FS(index))
}
