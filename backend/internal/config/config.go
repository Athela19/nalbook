package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port             string
	DatabaseAdminURL string
	TenantDSNTpl     string
	S3Endpoint       string
	S3Bucket         string
	JWTSecret        string
	FrontendOrigin   string
}

func Load() (Config, error) {
	c := Config{
		Port:             envOr("PORT", "8080"),
		DatabaseAdminURL: os.Getenv("DATABASE_ADMIN_URL"),
		TenantDSNTpl:     os.Getenv("TENANT_DSN_TPL"),
		S3Endpoint:       os.Getenv("S3_ENDPOINT"),
		S3Bucket:         envOr("S3_BUCKET", "nalbook"),
		JWTSecret:        os.Getenv("JWT_SECRET"),
		FrontendOrigin:   envOr("FRONTEND_ORIGIN", "http://localhost:3000"),
	}
	var missing []string
	if c.DatabaseAdminURL == "" {
		missing = append(missing, "DATABASE_ADMIN_URL")
	}
	if c.TenantDSNTpl == "" {
		missing = append(missing, "TENANT_DSN_TPL")
	}
	if c.JWTSecret == "" {
		missing = append(missing, "JWT_SECRET")
	}
	if len(missing) > 0 {
		return Config{}, fmt.Errorf("env wajib belum diisi: %v", missing)
	}
	return c, nil
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
