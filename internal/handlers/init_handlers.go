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

func Init() *echo.Echo {
	e := echo.New()
	e.Static("/", "static")
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	e.GET("/", routes.NewHomeHandler().Init)
	e.GET("/work", routes.NewWorkHandler().Init)

	return e
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	switch code {
	case 404:
		page := pages.NotFound()
		err = templates.Layout(page, "benmarshall").Render(context.Background(), c.Response().Writer)
		if err != nil {
			c.Logger().Error(err)
		}
	default:
		c.Logger().Error(err)
	}
}
