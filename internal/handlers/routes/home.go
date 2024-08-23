package routes

import (
	"context"

	"github.com/benleem/benmarshall/internal/templates"
	"github.com/benleem/benmarshall/internal/templates/pages"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Get(c echo.Context) error {
	page := pages.Home()
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
