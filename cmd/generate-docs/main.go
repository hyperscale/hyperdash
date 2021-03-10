// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"text/template"
)

var enumTempl = `Unit category
=============

see: https://github.com/grafana/grafana/blob/master/packages/grafana-data/src/valueFormats/categories.ts

{{range .Categories}}
{{ .Name }}
----

| Name  | Unit  |
|---|---|{{range .Formats}}
| {{.Name}} | ` + "`{{.Unit}}`" + ` |{{end}}

{{end}}
`

type itemFormat struct {
	Name string `json:"name"`
	Unit string `json:"unit"`
}

type itemCategory struct {
	Name    string       `json:"name"`
	Formats []itemFormat `json:"formats"`
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) == 1 {
		log.Fatalf("%s [path to dest file]\n", os.Args[0])
	}

	content, err := os.ReadFile(dir + "/resources/unit-categories.json")
	if err != nil {
		log.Fatal(err)
	}

	var categories []itemCategory

	if err := json.Unmarshal(content, &categories); err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.New("enum").Parse(enumTempl))

	buf := bytes.NewBuffer(nil)

	if err := t.Execute(buf, struct {
		Categories []itemCategory
	}{
		Categories: categories,
	}); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(os.Args[2], buf.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
}
