package main

import (
	"fmt"
	"net/http"

	"go-app/app/router"
	"go-app/config"
	lr "go-app/util/logger"
)

func main() {
	appConfig := config.AppConfig()

	logger := lr.New(appConfig.Debug)

	appRouter := router.New()

	address := fmt.Sprintf(":%d", appConfig.Server.Port)

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConfig.Server.TimeoutRead,
		WriteTimeout: appConfig.Server.TimeoutWrite,
		IdleTimeout:  appConfig.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}
