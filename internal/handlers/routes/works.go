package routes

import (
	"context"
	"fmt"

	"github.com/a-h/templ"
	"github.com/benleem/benmarshall/internal/templates"
	"github.com/benleem/benmarshall/internal/templates/pages"
	"github.com/labstack/echo/v4"
)

type WorksHandler struct {
	Works []pages.WorkInfo
}

func NewWorksHandler() *WorksHandler {
	works := []pages.WorkInfo{
		{Id: "love-together", Name: "Love Together", Tags: []string{"TypeScript", "NextJs", "Tailwind"}, Live: "https://www.love-together.com/", Code: ""},
		{Id: "prattl", Name: "prattl", Tags: []string{"Go", "Python"}, Live: "", Code: "https://github.com/benleem/prattl"},
		{Id: "llm-notes", Name: "llmnotes", Tags: []string{"Rust", "React Native", "Surreal Db"}, Live: "", Code: "https://github.com/benleem/prattl"},
		{Id: "show-beam", Name: "ShowBeam", Tags: []string{"Go"}, Live: "", Code: "https://github.com/benleem/show_beam_client"},
		{Id: "goggle-earth", Name: "Goggle Earth", Tags: []string{"JavaScript", "React", "CSS", "ThreeJs"}, Live: "https://goggle-earth.netlify.app/", Code: "https://github.com/benleem/threejs-test"},
		// {Id: "chirp", Name: "Chirp", Tags: []string{"JavaScript", "NextJs", "CSS Modules"}, Live: "https://chirp-social.vercel.app/", Code: "https://github.com/benleem/chirp"},
	}

	return &WorksHandler{
		Works: works,
	}
}

func (h *WorksHandler) GetWorks(c echo.Context) error {
	var page templ.Component
	hxReq := c.Request().Header.Get("Hx-Request")
	if hxReq != "" {
		page = pages.Works(true, h.Works)
		c.Response().Header().Set(echo.HeaderVary, "Hx-Request")
		return page.Render(context.Background(), c.Response().Writer)
	}
	page = pages.Works(false, h.Works)
	return templates.Layout(page, "benmarshall - works").Render(context.Background(), c.Response().Writer)

}

func (h *WorksHandler) GetWork(c echo.Context) error {
	id := c.Param("id")
	var workInfo pages.WorkInfo
	for _, work := range h.Works {
		if work.Id == id {
			workInfo = work
		}
	}
	if workInfo.Id == "" {
		return echo.ErrNotFound
	}
	var page templ.Component

	hxReq := c.Request().Header.Get("Hx-Request")
	if hxReq != "" {
		page = pages.Work(true, id, workInfo)
		c.Response().Header().Set(echo.HeaderVary, "Hx-Request")
		return page.Render(context.Background(), c.Response().Writer)
	}
	page = pages.Work(false, id, workInfo)
	return templates.Layout(page, fmt.Sprintf("benmarshall - %s", workInfo.Name)).Render(context.Background(), c.Response().Writer)
}
