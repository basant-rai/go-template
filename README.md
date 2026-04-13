# Go Template — Backend API

Go template structure for microservice. Built with **Go + Gin**, running on **PostgreSQL**.

Supports auto migration with gomigrate https://pkg.go.dev/github.com/basant-rai/gomigrate.

---

## Table of Contents

- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
- [Environment Variables](#environment-variables)
- [Swagger Docs](#swagger-docs)

---

## Architecture

The project follows a Hexagonal (Ports and Adapters) architecture. Each domain module contains its own:

```
internal/{module}/
  handler.go     -- HTTP handlers and Swagger docs
  service.go     -- Business logic (Service interface + implementation)
  repository.go  -- Data access (Repository interface + implementation)
  model.go       -- Domain structs and request/response DTOs
```

```
project/
  cmd/server/main.go           -- Entry point, dependency wiring
  config/config.go             -- Environment-based configuration
  api/
    routes.go                  -- Top-level route registration
    v1/routes.go               -- v1 API routes with rate limiting
  internal/
    bootstrap/handlers.go      -- Handler container struct
    domain/                    -- Base UUID+timestamp struct
    platform/
      database/                -- PostgreSQL connection pool
      middleware/              -- JWT auth, CORS, request logger, rate limiter
    integration                -- External Client Integrations
      stripe/client.go         -- Stripe client wrapper
      google/client.go         -- Google client wrapper
    email/
      sender.go                -- EmailSender and provider dispatch
      templates.go             -- baseLayout and email templates
    user/                      -- Internal Business logic
      user/handler.go          -- HTTP handlers and Swagger docs
      user/model.go            -- Business logic (Service interface + implementation)
      user/service.go          -- Data access (Repository interface + implementation)
      user/repository.go       -- Domain structs and request/response DTOs
  migrations/                  -- Sequential SQL migrations
  docs/                        -- Generated Swagger docs
```

---

## Tech Stack

| Layer            | Technology                                        |
| ---------------- | ------------------------------------------------- |
| Language         | Go 1.22+                                          |
| HTTP Framework   | Gin v1.12                                         |
| Database Driver  | pgx/v5 (native PostgreSQL)                        |
| Database         | PostgreSQL 15+                                    |
| Payments         | Stripe                                            |
| Auth             | golang-jwt/jwt v5 (HS256, 24h expiry)             |
| Password Hashing | bcrypt (cost 10)                                  |
| Email            | SMTP (gomail.v2)         |
| Rate Limiting    | In-memory sliding window (no external dependency) |
| API Docs         | Swagger via swaggo/gin-swagger                    |
| Live Reload      | Air                                               |
| Config           | godotenv + os.Getenv                              |

---

## Getting Started

### Prerequisites

```
go version      # requires Go 1.22+
psql --version  # PostgreSQL 15+
```

### Install

```bash
git clone <repo-url>
cd go-template
go mod download
go mod tidy
```

### Environment Setup

```bash
cp .env.example .env
# Edit .env with your values
```

### Run Migrations

All migration commands are available via Makefile:

```bash
# Install migrate CLI
make migrate-install

# Initialize migration
make migrate-init

# Create a new migration
make migrate-create NAME=add_new_table

# Check new migration
make migrate-diff

# Auto generate migration
make migrate-generate NAME=fuel_transaction_uuid_change;

# Apply all pending migrations
make migrate-up

# Rollback one migration
make migrate-down

# Show current migration version
make migrate-status

# Reset database (dev only)
make migrate-reset

# Build and run server
make dev
```

### Generate Swagger Docs

```bash
swag init -g cmd/server/main.go --output docs --parseDependency --parseInternal
```

### Run (with live reload)

```bash
go install github.com/air-verse/air@latest
air
```

### Run (without live reload)

```bash
go run ./cmd/server/main.go
```

Server starts on `http://localhost:8080` by default.

---

## Environment Variables

### Required

| Variable                 | Description                                        |
| ------------------------ | -------------------------------------------------- |
| `DATABASE_URL`           | PostgreSQL connection string                       |
| `JWT_SECRET`             | Secret for HS256 JWT signing                       |
| `STRIPE_SECRET_KEY`      | Stripe secret key — `sk_test_...` or `sk_live_...` |
| `STRIPE_PUBLISHABLE_KEY` | Stripe publishable key                             |

### Optional

| Variable        | Default                 | Description                             |
| --------------- | ----------------------- | --------------------------------------- |
| `PORT`          | `8080`                  | HTTP server port                        |
| `APP_ENV`       | `development`           | `development` or `production`           |
| `APP_URL`       | `http://localhost:3000` | Frontend URL for invite/reset links     |
| `CORS_ORIGINS`  | See below               | Comma-separated list of allowed origins |
| `EMAIL_FROM`    | `noreply@domain.com` | Sender email address                    |
| `SMTP_HOST`     | —                       | SMTP hostname. Priority 3.              |
| `SMTP_PORT`     | `587`                   | SMTP port                               |
| `SMTP_USER`     | —                       | SMTP username                           |
| `SMTP_PASSWORD` | —                       | SMTP password                           |

### CORS Origins

Set `CORS_ORIGINS` as a comma-separated string to configure allowed origins:

```
CORS_ORIGINS=https://domain.com
```

If not set, defaults to:

```
http://localhost:3000, https://domain.com
```

## Swagger Docs

Regenerate after any handler or model change:

```bash
swag init -g cmd/server/main.go --output docs --parseDependency --parseInternal
```

Access at: `http://localhost:8080/swagger/index.html`

