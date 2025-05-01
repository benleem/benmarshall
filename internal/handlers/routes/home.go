package routes

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/benleem/benmarshall/internal/templates"
	"github.com/benleem/benmarshall/internal/templates/pages"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context().Value("htmx"))

	var page templ.Component
	page = pages.Home()
	templates.Layout(page, "benmarshall").Render(r.Context(), w)
}
