package routes

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/benleem/benmarshall/internal/templates"
	"github.com/benleem/benmarshall/internal/templates/components"
	"github.com/benleem/benmarshall/internal/templates/pages"
)

type HomeHandler struct {
	skills      components.Skills
	skillKeys   []string
	projects    components.Projects
	projectKeys []string
}

func NewHomeHandler() *HomeHandler {
	skillKeys := []string{"tools", "languages", "technologies"}
	skills := components.Skills{
		skillKeys[0]: []string{"git"},
		skillKeys[1]: []string{"golang", "typescript", "javascript"},
		skillKeys[2]: []string{"docker", "html", "css", "tailwind", "react", "nextjs", "htmx"},
	}

	projectKeys := []string{"experience", "projects"}
	projects := components.Projects{
		projectKeys[0]: {{Name: `love together`, Description: "relationship coaching web application", Tags: []string{"typescript", "nextjs", "mongodb", "tailwind"}, Github: "", Live: "https://www.love-together.com/", Image: "https://res.cloudinary.com/ben-dev/image/upload/v1677532380/portfolio/love-together_zkbqt3.png"}},
		projectKeys[1]: {
			{Name: "prattl", Description: "cli tool for local audio transcriptions", Tags: []string{"go", "python"}, Github: "https://github.com/prattlOrg/prattl", Live: "https://prattl.co/", Image: "/static/assets/prattl.png"},
			{Name: "goggle earth", Description: "web application for learning about earth, place a pin and see information about the area", Tags: []string{"javascript", "react", "threejs"}, Github: "https://github.com/benleem/threejs-test", Live: "https://goggle-earth.netlify.app/", Image: "https://res.cloudinary.com/ben-dev/image/upload/v1677537530/portfolio/earth_zn3y1v.png"},
			{Name: "3ohtwo (WIP)", Description: "mobile application for marking your favorite places, and seeing other people's", Tags: []string{"typescript", "react native expo", "sqlite"}, Github: "https://github.com/benleem/3ohtwo", Live: "", Image: ""},
			{Name: "goscrape (WIP)", Description: "web crawling and scraping cli tool", Tags: []string{"go", "redis", "sqlite"}, Github: "", Live: "", Image: ""},
		},
	}

	return &HomeHandler{skills, skillKeys, projects, projectKeys}
}

func (h *HomeHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context().Value("htmx"))

	var page templ.Component
	page = pages.Home(h.skills, h.skillKeys, h.projects, h.projectKeys)
	templates.Layout(page, "benmarshall").Render(r.Context(), w)
}
