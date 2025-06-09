package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"moria-gate-api/internal/core/services"
	"moria-gate-api/internal/infra/db"
	"moria-gate-api/internal/infra/http/handler"
	"moria-gate-api/internal/infra/http/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	logger, _ := zap.NewProduction()
	r.Use(middleware.LoggerMiddleware(logger))
	r.Use(gin.Recovery())

	// DB
	connStr := os.Getenv("DB_DSN")
	repo, err := db.NewPostgresRepository(connStr)
	if err != nil {
		panic(err)
	}

	// Services
	lookup := services.NewLookupService(repo)
	secure := handler.NewSecureHandler(lookup)

	// Routes
	v1 := r.Group("/v1")
	v1.POST("/secure-data", secure.SecureData)

	return r
}
