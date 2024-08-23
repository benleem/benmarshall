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
	key, keyOk := os.LookupEnv("WEB3FORMSKEY")
	if !keyOk {
		log.Fatal("Web3Forms api key not defined")
	}

	e := handlers.Init(key)
	fmt.Println("✅ server running")
	fmt.Printf("localhost%s\n", port)
	e.Logger.Fatal(e.Start(port))
}
