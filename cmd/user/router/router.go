package router

import (
	"github.com/gin-gonic/gin"
	user_router "go_python/internal/app/user/router"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	apiGroup := router.Group("/v1")
	user_router.InitUserRouter(apiGroup)
	return router
}
