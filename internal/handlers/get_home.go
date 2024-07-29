package handlers

import (
	"context"

	"github.com/benleem/benmarshall/internal/templates"
	"github.com/benleem/benmarshall/internal/templates/components"
	"github.com/benleem/benmarshall/internal/templates/pages"
	"github.com/labstack/echo"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Init(c echo.Context) error {
	navData := []components.NavData{{"Work", "/work"}, {"Contact", "/contact"}}
	page := pages.Home(navData)
	dummyNav := []components.NavData{}
	return templates.Layout(page, "benmarshall", dummyNav).Render(context.Background(), c.Response().Writer)
	// if err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }

}
