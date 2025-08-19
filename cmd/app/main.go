package main

import (
	"fmt"
	"log"
	"os"
	"usr-short/cmd/internal/config"
	http_request_middleware "usr-short/cmd/internal/http-server/middleware"
	"usr-short/cmd/internal/logger"
	"usr-short/cmd/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error uploading file local.env", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error uploading config", err)
	}

	log := logger.SetupLogger(cfg.ENV)

	_, err = storage.Start(cfg.STORAGE_PATH)
	if err != nil {
		log.Error("error to init storage", err)
		os.Exit(1)
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(http_request_middleware.HttpRequestTrace(log))

	fmt.Println(cfg)
}
