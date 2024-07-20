package handler

import (
	"benmarshall/render"
	"log"
	"net/http"
)

type HomeData struct {
	Message string
}

func Home(w http.ResponseWriter, r *http.Request, tmpls *render.Templates) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := tmpls.Render(w, "home", HomeData{"Home"})
	if err != nil {
		log.Println("template render error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type WorkData struct {
	Message string
}

func Work(w http.ResponseWriter, r *http.Request, tmpls *render.Templates) {
	err := tmpls.Render(w, "work", WorkData{"Work"})
	if err != nil {
		log.Println("template render error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
