package main

import (
	"fmt"
	"log"
	"net/http"

	"go-app/app/router"
	"go-app/config"
)

func main() {
	appConfig := config.AppConfig()

	appRouter := router.New()

	address := fmt.Sprintf(":%d", appConfig.Server.Port)

	log.Println("Running on server 6969")

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConfig.Server.TimeoutRead,
		WriteTimeout: appConfig.Server.TimeoutWrite,
		IdleTimeout:  appConfig.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
