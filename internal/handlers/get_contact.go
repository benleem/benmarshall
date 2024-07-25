package handlers

import (
	"benmarshall/internal/templates"
	"context"

	"github.com/labstack/echo"
)

type ContactHandler struct{}

func NewContactHandler() *ContactHandler {
	return &ContactHandler{}
}

func (h *ContactHandler) Init(c echo.Context) error {
	page := templates.Contact()
	return templates.Layout(page, "benmarshall").Render(context.Background(), c.Response().Writer)
	// if err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }
}
