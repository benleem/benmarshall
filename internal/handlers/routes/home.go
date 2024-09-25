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
	hxReq := c.Request().Header.Get("Hx-Request")
	if hxReq != "" {
		c.Response().Header().Set(echo.HeaderVary, "Hx-Request")
		return page.Render(context.Background(), c.Response().Writer)
	}
	return templates.Layout(page, "benmarshall").Render(context.Background(), c.Response().Writer)
}
