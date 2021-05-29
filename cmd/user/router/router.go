package router

import (
	"github.com/gin-gonic/gin"
	user_router "go_python/internal/app/user/router"
	"go_python/internal/pkg/logging"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(logging.GinZapLogger(), gin.Recovery())
	userApiGroup := router.Group("/u/v1")
	user_router.InitUserRouter(userApiGroup)
	return router
}
