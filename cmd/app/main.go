package main

import (
	"fmt"
	"log"
	"usr-short/cmd/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error uploading file local.env:", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error uploading config:", err)
	}

	fmt.Println(cfg)
}
