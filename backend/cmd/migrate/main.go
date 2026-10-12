package main

import (
	"log/slog"
	"os"

	"nalbook/backend/internal/config"
)

func main() {
	if _, err := config.Load(); err != nil {
		slog.Error("konfigurasi tidak valid", "error", err)
		os.Exit(1)
	}
	slog.Info("migrasi selesai")
}
