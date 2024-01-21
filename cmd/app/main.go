package main

import (
	"log"
	"net/http"
	"os"

	"caching-user-app/config"
	"caching-user-app/pkg/database"
	"caching-user-app/pkg/routes"
)

func main() {
	configApp, errConfig := config.LoadConfig()

	if errConfig != nil {
		log.Fatal("Loading config failed: ", errConfig)
		os.Exit(1)
	}

	_, errDb := database.Init(configApp.DbUrl)

	if errDb != nil {
		log.Fatal("Cannot connect to db: ", errDb)
		os.Exit(1)
	}

	srv := &http.Server{
		Handler: routes.NewRouter(),
		Addr:    ":" + configApp.Port,
	}

	log.Printf("Server starting on port %v", configApp.Port)
	errServer := srv.ListenAndServe()

	if errServer != nil {
		log.Fatal(errServer)
	}

	defer srv.Close()
}
