package main

import (
	t "text/template"
	h "html/template"
)

func main() {
	t.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	h.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
}