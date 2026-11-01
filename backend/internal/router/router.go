package router

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"stockai/backend/internal/handler"
)

func New(db *gorm.DB, redisClient *redis.Client) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())

	healthHandler := handler.NewHealthHandler(db, redisClient)

	api := engine.Group("/api/v1")
	api.GET("/health", healthHandler.Check)

	return engine
}
