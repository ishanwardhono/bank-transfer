package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ishanwardhono/transfer-system/config"
	"github.com/ishanwardhono/transfer-system/pkg/logger"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("[MAIN] Failed to load config: %v", err)
	}
	logger.Init(cfg.LogLevel)

	appContainer, err := NewAppContainer(cfg)
	if err != nil {
		logger.Fatalf(ctx, "[MAIN] Failed to initialize application: %v", err)
	}

	// Create channel for shutdown signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Run server in a goroutine
	serverErrors := make(chan error, 1)
	go func() {
		logger.Printf(ctx, "[MAIN] Starting app server on %s", cfg.GetServerAddress())
		err := appContainer.Run()
		if err != nil && err != http.ErrServerClosed {
			// Only report errors that aren't related to shutdown
			serverErrors <- err
		}
	}()

	// Wait for shutdown signal or server error
	select {
	case <-stop:
		logger.Info(ctx, "[MAIN] Received shutdown signal. Server stopping...")
	case err := <-serverErrors:
		logger.Errorf(ctx, "[MAIN] Server error: %v", err)
	}

	// Create shutdown timeout context
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Shutdown server
	logger.Info(ctx, "[MAIN] Initiating graceful shutdown...")
	if err := appContainer.Shutdown(ctx); err != nil {
		logger.Errorf(ctx, "[MAIN] Server shutdown failed: %v", err)
		os.Exit(1)
	}

	logger.Info(ctx, "[MAIN] Server stopped gracefully")
}
