package main

import (
	"benmarshall/internal/handlers"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	_ = godotenv.Load(".env")
	port, portOk := os.LookupEnv("PORT")
	if !portOk {
		log.Fatal("Port not defined")
	}

	e := echo.New()
	e.Static("/", "static")
	e.GET("/", handlers.NewHomeHandler().Init)
	e.GET("/work", handlers.NewWorkHandler().Init)
	e.GET("/contact", handlers.NewContactHandler().Init)

	fmt.Println("✅ server running")
	fmt.Printf("localhost%s\n", port)
	e.Logger.Fatal(e.Start(port))

	// mux := http.NewServeMux()
	// mux.Handle("/public/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.Home(w, r, tmpls)
	// })
	// mux.HandleFunc("/work/", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.Work(w, r, tmpls)
	// })
	// mux.HandleFunc("/contact/", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.Contact(w, r, tmpls)
	// })
	// log.Fatal(http.ListenAndServe(port, mux))

}
