package render

import (
	"io"
	"text/template"
)

type Template struct {
	tmpl *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.tmpl.ExecuteTemplate(w, name+".html", data)
}

func NewTemplateRenderer(path string) *Template {
	return &Template{
		tmpl: template.Must(template.ParseGlob(path)),
	}
}
