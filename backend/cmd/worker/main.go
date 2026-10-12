package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"nalbook/backend/internal/config"
)

func main() {
	if _, err := config.Load(); err != nil {
		slog.Error("konfigurasi tidak valid", "error", err)
		os.Exit(1)
	}
	slog.Info("worker nalbook siap")
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
