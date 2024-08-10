package handlers

import (
	"context"
	"net/http"

	"github.com/benleem/benmarshall/internal/handlers/routes"
	"github.com/benleem/benmarshall/internal/templates"
	"github.com/benleem/benmarshall/internal/templates/components"
	"github.com/benleem/benmarshall/internal/templates/pages"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Static("/", "static")
	e.GET("/", routes.NewHomeHandler().Init)
	e.GET("/work", routes.NewWorkHandler().Init)
	e.GET("/contact", routes.NewContactHandler().Init)

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
		navData := []components.NavData{{"Home", "/"}, {"Work", "/work"}, {"Contact", "/contact"}}
		err = templates.Layout(page, "benmarshall", navData).Render(context.Background(), c.Response().Writer)
		if err != nil {
			c.Logger().Error(err)
		}
	default:
		c.Logger().Error(err)
	}
}
