package routes

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/benleem/benmarshall/internal/templates"
	"github.com/benleem/benmarshall/internal/templates/pages"
)

type HomeHandler struct {
	skills      pages.Skills
	skillKeys   []string
	projects    pages.Projects
	projectKeys []string
}

func NewHomeHandler() *HomeHandler {
	skillKeys := []string{"tools", "languages", "technologies"}
	skills := pages.Skills{
		skillKeys[0]: []string{"git"},
		skillKeys[1]: []string{"golang", "typescript", "javascript"},
		skillKeys[2]: []string{"docker", "html", "css", "tailwind", "react", "nextjs", "htmx"},
	}

	projectKeys := []string{"experience", "projects"}
	projects := pages.Projects{
		projectKeys[0]: {{Name: "Love Together", Description: " Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras at tortor mauris. Suspendisse at dolor ac eros iaculis tristique eget tincidunt quam. Suspendisse eu accumsan sapien. Suspendisse potenti. Sed auctor massa diam, at sodales mi gravida non. Etiam a vestibulum sem, at iaculis odio. Sed ut velit vitae risus suscipit interdum a nec dui. Nam pellentesque interdum fringilla. ", Tags: []string{"asdad", "sadasd"}, Github: "", Live: "", Image: ""}},
		projectKeys[1]: {
			{Name: "prattl", Description: "ipsum dolor sit amet, consectetur adipiscing elit. Cras at tortor mauris. Sed ut velit vitae risus suscipit interdum a nec dui. Nam pellentesque interdum fringilla. ", Tags: []string{"asdad", "sadasd"}, Github: "", Live: "", Image: ""},
			{Name: "3ohtwo", Description: "Cras at tortor mauris. Suspendisse at dolor ac eros iaculis tristique eget tincidunt quam. Suspendisse eu accumsan sapien. Suspendisse potenti. Sed auctor massa diam, at sodales mi gravida non. Etiam a vestibulum sem, at iaculis odio. Sed ut velit vitae risus suscipit interdum a nec dui. Nam pellentesque interdum fringilla. ", Tags: []string{"asdad", "sadasd"}, Github: "", Live: "", Image: ""},
			{Name: "goscrape", Description: " Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras at tortor mauris. Suspendisse at dolor ac eros iaculis tristique eget tincidunt quam. Suspendisse eu accumsan sapien. Suspendisse potenti. Sed auctor massa diam, ", Tags: []string{"asdad", "sadasd"}, Github: "", Live: "", Image: ""},
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
