package render

import (
	"io"
	"text/template"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.Templates.ExecuteTemplate(w, name+".html", data)
}

func NewTemplateRenderer(paths ...string) *template.Template {
	tmpl := &template.Template{}
	for i := range paths {
		template.Must(tmpl.ParseGlob(paths[i]))
	}
	return tmpl
}
