package main

import (
	"benmarshall/handler"
	"benmarshall/render"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	port, portOk := os.LookupEnv("PORT")
	if !portOk {
		log.Fatal("Port not defined")
	}

	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("client/public/"))))
	tmpl := render.NewTemplateRenderer("client/public/html/*.html")
	renderer := &render.Template{Templates: tmpl}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.Home(w, r, renderer)
	})

	fmt.Println("✅ server running")
	fmt.Printf("localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
