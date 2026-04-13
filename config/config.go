package config

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                 string
	AppEnv               string
	DatabaseURL          string
	RedisURL             string
	JWTSecret            string
	StripeSecretKey      string
	StripePublishableKey string
	StripeWebhookSecret  string
	BrevoAPIKey          string
	SMTPHost             string
	SMTPPort             string
	SMTPUser             string
	SMTPPassword         string
	ResendAPIKey         string
	EmailFrom            string
	AppURL               string
	CORSOrigins          []string
	EncryptionKey        []byte
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("[config] no .env file — using environment variables")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}
	stripeSecret := os.Getenv("STRIPE_SECRET_KEY")
	if stripeSecret == "" {
		return nil, fmt.Errorf("STRIPE_SECRET_KEY is required")
	}
	stripePub := os.Getenv("STRIPE_PUBLISHABLE_KEY")
	if stripePub == "" {
		return nil, fmt.Errorf("STRIPE_PUBLISHABLE_KEY is required")
	}

	encryptionKeyHex := os.Getenv("ENCRYPTION_KEY")
	if encryptionKeyHex == "" {
		return nil, fmt.Errorf("ENCRYPTION_KEY is required (32-byte hex string for AES-256)")
	}
	encryptionKey, err := hex.DecodeString(encryptionKeyHex)
	if err != nil {
		return nil, fmt.Errorf("ENCRYPTION_KEY must be valid hex: %w", err)
	}
	if len(encryptionKey) != 32 {
		return nil, fmt.Errorf("ENCRYPTION_KEY must be 32 bytes (64 hex chars), got %d bytes", len(encryptionKey))
	}

	return &Config{
		Port:                 getEnv("PORT", "8080"),
		AppEnv:               getEnv("APP_ENV", "development"),
		DatabaseURL:          dbURL,
		RedisURL:             getEnv("REDIS_URL", "redis://localhost:6379"),
		JWTSecret:            jwtSecret,
		StripeSecretKey:      stripeSecret,
		StripePublishableKey: stripePub,
		StripeWebhookSecret:  getEnv("STRIPE_WEBHOOK_SECRET", ""),
		BrevoAPIKey:          getEnv("BREVO_API_KEY", ""),
		SMTPHost:             getEnv("SMTP_HOST", ""),
		SMTPPort:             getEnv("SMTP_PORT", "587"),
		SMTPUser:             getEnv("SMTP_USER", ""),
		SMTPPassword:         getEnv("SMTP_PASSWORD", ""),
		ResendAPIKey:         getEnv("RESEND_API_KEY", ""),
		EmailFrom:            getEnv("EMAIL_FROM", "noreply@supertruck.ai"),
		AppURL:               getEnv("APP_URL", "http://localhost:3000"),
		CORSOrigins:          parseCORSOrigins(getEnv("CORS_ORIGINS", "")),
		EncryptionKey:        encryptionKey,
	}, nil
}

// parseCORSOrigins parses a comma-separated CORS_ORIGINS env var.
// Falls back to safe defaults when not set.
func parseCORSOrigins(raw string) []string {
	if raw != "" {
		var origins []string
		for _, o := range strings.Split(raw, ",") {
			if trimmed := strings.TrimSpace(o); trimmed != "" {
				origins = append(origins, trimmed)
			}
		}
		if len(origins) > 0 {
			return origins
		}
	}
	return []string{
		"http://localhost:3000",
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
