package main

import (
	"log"
	"net/http"
	"os"
	"usr-short/cmd/internal/config"
	http_request_middleware "usr-short/cmd/internal/http-server/middleware"
	"usr-short/cmd/internal/logger"
	service "usr-short/cmd/internal/service/url"
	"usr-short/cmd/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
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

	storage, err := storage.Start(cfg.STORAGE_PATH)
	if err != nil {
		log.Error("Error to init storage", err)
		os.Exit(1)
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(http_request_middleware.HttpRequestTrace(log))
	router.Use(middleware.Recoverer)

	router.Get("/url", service.GetUrl(log, storage))
	router.Post("/url", service.SaveUrl(log, storage))
	router.Delete("/url", service.DeleteUrl(log, storage))
	router.Put("/url", service.UpdateUrl(log, storage))

	server := &http.Server{
		Addr:         cfg.ADDRESS,
		Handler:      router,
		ReadTimeout:  cfg.TIMEOUT,
		WriteTimeout: cfg.IDLE_TIMEOUT,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Error("Error starting the server:", err)
	}

}
