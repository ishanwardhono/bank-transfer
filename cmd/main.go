package main

import (
	"context"
	"log"

	"github.com/ishanwardhono/transfer-system/config"
	"github.com/ishanwardhono/transfer-system/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	logger.Init(cfg.LogLevel)

	appContainer, err := NewAppContainer(cfg)
	if err != nil {
		logger.Fatalf(context.Background(), "Failed to initialize application: %v", err)
	}

	serverAddress := cfg.GetServerAddress()
	logger.Printf(context.Background(), "Starting HTTP server on %s", serverAddress)
	if err := appContainer.RunHTTPServer(serverAddress); err != nil {
		logger.Fatalf(context.Background(), "Failed to start HTTP server: %v", err)
	}
}
