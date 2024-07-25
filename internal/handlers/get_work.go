package handlers

import (
	"benmarshall/internal/templates"
	"context"

	"github.com/labstack/echo"
)

type WorkHandler struct{}

func NewWorkHandler() *WorkHandler {
	return &WorkHandler{}
}

func (h *WorkHandler) Init(c echo.Context) error {
	page := templates.Work()
	return templates.Layout(page, "benmarshall").Render(context.Background(), c.Response().Writer)
	// if err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }
}
