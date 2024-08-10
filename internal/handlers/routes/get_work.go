package routes

import (
	"context"

	"github.com/benleem/benmarshall/internal/templates"
	"github.com/benleem/benmarshall/internal/templates/components"
	"github.com/benleem/benmarshall/internal/templates/pages"
	"github.com/labstack/echo"
)

type WorkHandler struct{}

func NewWorkHandler() *WorkHandler {
	return &WorkHandler{}
}

func (h *WorkHandler) Init(c echo.Context) error {
	page := pages.Work()
	navData := []components.NavData{{"Home", "/"}, {"Contact", "/contact"}}
	return templates.Layout(page, "benmarshall", navData).Render(context.Background(), c.Response().Writer)
	// if err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }
}
