package routes

import (
	"context"

	"github.com/benleem/benmarshall/internal/templates"
	"github.com/benleem/benmarshall/internal/templates/components"
	"github.com/benleem/benmarshall/internal/templates/pages"
	"github.com/labstack/echo"
)

type ContactHandler struct{}

func NewContactHandler() *ContactHandler {
	return &ContactHandler{}
}

func (h *ContactHandler) Init(c echo.Context) error {
	page := pages.Contact()
	navData := []components.NavData{{"Home", "/"}, {"Work", "/work"}}
	return templates.Layout(page, "benmarshall", navData).Render(context.Background(), c.Response().Writer)
	// if err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }
}
