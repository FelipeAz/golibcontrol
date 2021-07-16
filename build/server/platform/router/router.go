package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/platform/router/build"
	"github.com/FelipeAz/golibcontrol/infra/jwt"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/infra/redis"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/handler"

	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

// Run Starts the server
func Run(dbService *service.MySQLService, cache *redis.Cache) error {
	return buildRoutes(dbService, cache)
}

func buildRoutes(dbService *service.MySQLService, cache *redis.Cache) error {
	router := gin.Default()
	jwtAuth := jwt.NewAuth(cache)
	tokenAuthMiddleware := middleware.NewTokenMiddleware(jwtAuth)

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	cHandler := handler.NewCommentHandler(dbService)
	build.CommentRoutes(tokenAuthMiddleware, vGroup, cHandler)

	return router.Run(":8083")
}
