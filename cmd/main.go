package main

import (
	"fmt"
	"log"
	"os"

	"github.com/benleem/benmarshall/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	port, portOk := os.LookupEnv("PORT")
	if !portOk {
		log.Fatal("Port not defined")
	}

	e := handlers.Init()
	fmt.Println("✅ server running")
	fmt.Printf("localhost%s\n", port)
	e.Logger.Fatal(e.Start(port))
}
