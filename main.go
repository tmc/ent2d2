package main

import (
	"flag"
	"log"
	"os"
	"text/template"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	flag.Parse()
	path := "./ent/schema"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	graph, err := entc.LoadGraph(path, &gen.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err := tmpl.Execute(os.Stdout, graph); err != nil {
		log.Fatal(err)
	}
}

var tmpl = template.Must(template.New("d2-diagram").
	Funcs(template.FuncMap{
		"fmtType": func(s string) string {
			return s
		},
	}).
	Parse(`
{{- with $.Nodes }}
{{- range $n := . }}
{{ $n.Name }}: {
	shape: sql_table
	{{- if $n.HasOneFieldID }}
	{{ $n.ID.Name }}: {{ fmtType $n.ID.Type.String }} {constraint: primary_key}
	{{- end }}
	{{- range $f := $n.Fields }}
	{{ $f.Name }}: {{ fmtType $f.Type.String }} {{ if $f.IsEdgeField }} {constraint: foreign_key}{{ end }}{{ if $f.Unique }}{constraint: unique}{{ end }}
	# this is a comment

	{{- end }}
}
{{- end }}

{{- range $n := . }}
    {{- range $e := $n.Edges }}
	{{- if not $e.IsInverse }}
		{{- $rt := "->" }}{{ if $e.O2M }}{{ $rt = "->" }}{{ else if $e.M2O }}{{ $rt = "<-" }}{{ else if $e.M2M }}{{ $rt = "<->" }}{{ end }}
    	{{ $n.Name }} {{ $rt }} {{ $e.Type.Name }} : "{{ $e.Name }}{{ with $e.Ref }}/{{ .Name }}{{ end }}" {
		{{ if $e.O2O }}
		source-arrowhead: {
			shape: cf-one
		}
		target-arrowhead: {
			shape: {{ if $e.Optional }}cf-one{{ else }}cf-one-required{{ end }}
		}
		{{ else if $e.O2M }}
		source-arrowhead: {
			shape: cf-one-required
		}
		target-arrowhead: * {
			shape: cf-many
		}
		{{ else if $e.M2O }}
		source-arrowhead: {
			shape: cf-many
		}
		target-arrowhead: {
			shape: cf-one
		}
		{{ else if $e.M2M }}
		source-arrowhead: * {
			shape: cf-many
		}
		target-arrowhead: * {
			shape: {{ if $e.Required }}cf-many-required{{ else }}cf-many{{ end }}
		}
		{{ end }}
	}
	{{- end }}
	{{- end }}
{{- end }}

{{- end }}
`))
