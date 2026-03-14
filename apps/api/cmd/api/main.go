package main

import (
	"log"
	"net/http"

	"WebParkir/apps/api/internal/delivery/http/handler"
	"WebParkir/apps/api/internal/infrastructure/config"
	"WebParkir/apps/api/internal/infrastructure/database"
)

func main() {

	// load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// connect database
	db, err := database.ConnectPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("database connected")

	// router
	mux := http.NewServeMux()

	// health endpoint
	mux.HandleFunc("/health", handler.HealthHandler)

	// start server
	addr := ":" + cfg.Port

	log.Println("server running on", addr)

	err = http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
