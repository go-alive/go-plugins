package plugin

var (
	tmpl = `
package main

import (
	"{{.Path}}"
	"github.com/go-alive/go-plugins"
)

var Plugin = plugin.Plugin{
	Name: "{{.Name}}",
	Type: "{{.Type}}",
	Path: "{{.Path}}",
	NewFunc: {{.Name}}.{{.NewFunc}},
}
`
)
