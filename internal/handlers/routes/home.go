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
		projectKeys[0]: {{Name: `love together`, Description: "relationship coaching web application", Tags: []string{"TypeScript", "NextJs", "Tailwind"}, Github: "", Live: "https://www.love-together.com/", Image: "https://res.cloudinary.com/ben-dev/image/upload/v1677532380/portfolio/love-together_zkbqt3.png"}},
		projectKeys[1]: {
			{Name: "prattl", Description: "cli tool for local transcriptions", Tags: []string{"Go", "Python"}, Github: "https://github.com/prattlOrg/prattl", Live: "https://prattl.co/", Image: "https://github.com/prattlOrg/prattl/blob/main/assets/logo.png"},
			{Name: "3ohtwo", Description: "mark your favorite places, see other people's", Tags: []string{"Typescript", "React Native Expo", "SqlLite"}, Github: "https://github.com/benleem/3ohtwo", Live: "", Image: ""},
			{Name: "goscrape", Description: "web crawling and scraping cli tool", Tags: []string{"Go", "Redis", "SqlLite"}, Github: "", Live: "", Image: ""},
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
