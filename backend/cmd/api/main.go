package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"nalbook/backend/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("konfigurasi tidak valid", "error", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/api/v1/health", healthz)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           withTraceID(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	slog.Info("backend nalbook berjalan", "port", cfg.Port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("server berhenti", "error", err)
		os.Exit(1)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func withTraceID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Trace-ID", time.Now().Format("20060102150405"))
		next.ServeHTTP(w, r)
	})
}
