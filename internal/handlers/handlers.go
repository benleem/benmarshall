package handlers

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/benleem/benmarshall/internal/handlers/routes"
	"github.com/benleem/benmarshall/internal/templates"
	"github.com/benleem/benmarshall/internal/templates/pages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func Init(key string) *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Static("/", "static")
	home := routes.NewHomeHandler()
	works := routes.NewWorksHandler()
	contact := routes.NewContactHandler(key)

	e.GET("/", home.Get)
	e.GET("/works", works.GetWorks)
	e.GET("/works/:id", works.GetWork)
	e.GET("/contact", contact.Get)
	e.POST("/contact", contact.Post)

	return e
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	switch code {
	case 404:
		var page templ.Component
		hxReq := c.Request().Header.Get("Hx-Request")
		if hxReq != "" {
			page = pages.NotFound(true)
			c.Response().Header().Set(echo.HeaderVary, "Hx-Request")
			err := page.Render(context.Background(), c.Response().Writer)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			page = pages.NotFound(false)
			err := templates.Layout(page, "benmarshall - 404").Render(context.Background(), c.Response().Writer)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	default:
		c.Logger().Error(err)
	}
}

// func middleware(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		hxReq := r.Header.Get("Hx-Request")

// 		if hxReq == "" {
// 			urlParts := strings.Split(r.URL.String(), "?")
// 			path := urlParts[0]
// 			params := ""

// 			if len(urlParts) > 1 {
// 				params = string('?') + urlParts[1]
// 			}

// 			pageVariables := PageVariables{
// 				Path:   path,
// 				Params: params,
// 			}

// 			tmpl := template.Must(template.ParseFiles(IndexTemplates...))

// 			err := tmpl.Execute(w, pageVariables)
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 				return
// 			}
// 			return
// 		}
// 		h.ServeHTTP(w, r)
// 	})
// }
