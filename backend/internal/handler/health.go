package handler

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const (
	serviceName       = "stockai-backend"
	appVersion        = "0.1.0"
	dependencyTimeout = 2 * time.Second
)

type HealthHandler struct {
	db    *gorm.DB
	redis *redis.Client
}

type healthResponse struct {
	Status       string            `json:"status"`
	Service      string            `json:"service"`
	Version      string            `json:"version"`
	Timestamp    time.Time         `json:"timestamp"`
	Dependencies map[string]string `json:"dependencies"`
}

func NewHealthHandler(db *gorm.DB, redisClient *redis.Client) *HealthHandler {
	return &HealthHandler{
		db:    db,
		redis: redisClient,
	}
}

func (h *HealthHandler) Check(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), dependencyTimeout)
	defer cancel()

	dependencies := map[string]string{
		"mysql": "ok",
		"redis": "ok",
	}
	status := "ok"

	sqlDB, err := h.db.DB()
	if err == nil {
		if pingErr := sqlDB.PingContext(ctx); pingErr != nil {
			dependencies["mysql"] = "down"
			status = "degraded"
			slog.WarnContext(ctx, "mysql health check failed", "error", pingErr)
		}
	} else {
		dependencies["mysql"] = "down"
		status = "degraded"
		slog.WarnContext(ctx, "mysql connection pool unavailable", "error", err)
	}

	if pingErr := h.redis.Ping(ctx).Err(); pingErr != nil {
		dependencies["redis"] = "down"
		status = "degraded"
		slog.WarnContext(ctx, "redis health check failed", "error", pingErr)
	}

	httpStatus := http.StatusOK
	if status != "ok" {
		httpStatus = http.StatusServiceUnavailable
	}

	c.JSON(httpStatus, healthResponse{
		Status:       status,
		Service:      serviceName,
		Version:      appVersion,
		Timestamp:    time.Now(),
		Dependencies: dependencies,
	})
}
