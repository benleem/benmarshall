package handlers

import (
	"benmarshall/internal/templates"
	"context"

	"github.com/labstack/echo"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Init(c echo.Context) error {
	page := templates.Home()
	return templates.Layout(page, "benmarshall").Render(context.Background(), c.Response().Writer)
	// if err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }
}
