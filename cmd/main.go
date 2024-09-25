package main

import (
	"fmt"
	"log"
	"os"

	"github.com/benleem/benmarshall/internal/handlers"
	"github.com/joho/godotenv"
)

type Config struct {
	port     string
	emailKey string
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load(".env")
	port, portOk := os.LookupEnv("PORT")
	if !portOk {
		return nil, fmt.Errorf("port not defined")
	}
	port = fmt.Sprintf(":%v", port)
	emailKey, keyOk := os.LookupEnv("WEB3FORMSKEY")
	if !keyOk {
		return nil, fmt.Errorf("web3forms api key not defined")
	}
	return &Config{
		port,
		emailKey,
	}, nil
}

func main() {
	config, err := NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	e := handlers.Init(config.emailKey)
	fmt.Println("✅ server running")
	fmt.Printf("localhost%s\n", config.port)
	e.Logger.Fatal(e.Start(config.port))
}
