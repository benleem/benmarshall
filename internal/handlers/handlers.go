package handlers

import (
	"context"
	"net/http"

	"github.com/benleem/benmarshall/internal/handlers/routes"
	"github.com/benleem/benmarshall/internal/templates"
	"github.com/benleem/benmarshall/internal/templates/pages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func Init(key string) *echo.Echo {
	e := echo.New()

	// e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Static("/", "static")
	home := routes.NewHomeHandler()
	work := routes.NewWorkHandler()
	contact := routes.NewContactHandler(key)

	e.GET("/", home.Get)
	e.GET("/work", work.Get)
	e.GET("/contact", contact.Get)
	e.POST("/contact", contact.Post)

	return e
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	switch code {
	case 404:
		page := pages.NotFound()
		hxReq := c.Request().Header.Get("Hx-Request")
		if hxReq != "" {
			c.Response().Header().Set(echo.HeaderVary, "Hx-Request")
			err := page.Render(context.Background(), c.Response().Writer)
			if err != nil {
				c.Logger().Error(err)
			}
		}
		err := templates.Layout(page, "benmarshall - 404").Render(context.Background(), c.Response().Writer)
		if err != nil {
			c.Logger().Error(err)
		}
	}
}
