package routes

import (
	"context"

	"github.com/benleem/benmarshall/internal/templates"
	"github.com/benleem/benmarshall/internal/templates/pages"
	"github.com/labstack/echo/v4"
)

type WorkHandler struct{}

func NewWorkHandler() *WorkHandler {
	return &WorkHandler{}
}

func (h *WorkHandler) Init(c echo.Context) error {
	works := []string{"Love Together", "prattl", "Goggle Earth", "Chirp", "Landing Page"}
	page := pages.Work(works)
	hxReq := c.Request().Header.Get("hx-request")
	if hxReq != "" {
		return page.Render(context.Background(), c.Response().Writer)
	}
	return templates.Layout(page, "benmarshall").Render(context.Background(), c.Response().Writer)
	// if err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }
}
