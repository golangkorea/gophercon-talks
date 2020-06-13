package main

import (
	h "html/template"
	t "text/template"
)

func main() {
	t.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	h.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
}
