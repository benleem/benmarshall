package handler

import (
	"benmarshall/render"
	"log"
	"net/http"
)

type HomeData struct {
	Message string
}

func Home(w http.ResponseWriter, r *http.Request, renderer *render.Template) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := renderer.Render(w, "index", HomeData{"Hello, world"})
	if err != nil {
		log.Println("template render error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
