package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	v1 "github.com/supertruck/wallet/api/v1"
	"github.com/supertruck/wallet/internal/bootstrap"
)

func RegisterRoutes(r *gin.Engine, h bootstrap.Handlers, jwtSecret string, db *pgxpool.Pool) {
	r.GET("/health", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()

		if err := db.Ping(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":  "unavailable",
				"service": "supertruck-wallet",
				"error":   "database connection failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "supertruck-wallet",
		})
	})

	//* ** WEBHOOKS ***
	r.POST("/webhooks/stripe", h.Webhook.HandleWebhook)

	//* ** SWAGGER ***
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//* ** API ***
	v1.Register(r, h, jwtSecret)
	// v2.Register(r, h, jwtSecret)

}
