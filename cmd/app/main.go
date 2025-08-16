package main

import (
	"fmt"
	"log"
	"usr-short/cmd/internal/config"
	"usr-short/cmd/internal/logger"

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

	log := logger.SetupLogger(cfg.ENV)

	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")

	fmt.Println(cfg)
}
