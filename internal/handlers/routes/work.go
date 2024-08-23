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

func (h *WorkHandler) Get(c echo.Context) error {
	works := []pages.WorkInfo{
		{Name: "Love Together", Tags: []string{"TypeScript", "NextJs", "Tailwind"}, Live: "https://www.love-together.com/", Code: ""},
		{Name: "prattl", Tags: []string{"Go", "Python"}, Live: "", Code: "https://github.com/benleem/prattl"},
		{Name: "ShowBeam (client)", Tags: []string{"TypeScript", "SvelteKit", "Tailwind"}, Live: "", Code: "https://github.com/benleem/show_beam_client"},
		{Name: "ShowBeam (server)", Tags: []string{"Rust", "Actix Web", "MySQL", "OAuth 2.0"}, Live: "", Code: "https://github.com/benleem/show_beam_server"},
		{Name: "Goggle Earth", Tags: []string{"JavaScript", "React", "CSS", "ThreeJs"}, Live: "https://goggle-earth.netlify.app/", Code: "https://github.com/benleem/threejs-test"},
		{Name: "Chirp", Tags: []string{"JavaScript", "NextJs", "CSS Modules"}, Live: "https://chirp-social.vercel.app/", Code: "https://github.com/benleem/chirp"},
		{Name: "CanvaCast", Tags: []string{"HTML", "JavaScript", "CSS"}, Live: "https://htmlpreview.github.io/?https://github.com/benleem/CanvaCast/blob/main/index.html", Code: "https://github.com/benleem/CanvaCast"},
	}
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
