package main

import (
	"fmt"
	"log"
	"os"
	"usr-short/cmd/internal/config"
	"usr-short/cmd/internal/logger"
	"usr-short/cmd/internal/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error uploading file local.env", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error uploading config", err)
	}

	log := logger.SetupLogger(cfg.ENV)

	storage, err := storage.Start(cfg.STORAGE_PATH)
	if err != nil {
		log.Error("Error to init storage", err)
		os.Exit(1)
	}
	res, err := storage.GetUrl("kexlet")
	fmt.Println(res)
	if err != nil {
		log.Error("Error to add url", err)

	}

	fmt.Println(cfg)
}
