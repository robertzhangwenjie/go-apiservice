package router

import (
	"github.com/gin-gonic/gin"
	"go_python/internal/app/user/api"
)

// 注册userGroup到router中
func InitUserRouter(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("list", api.GetUserList)
	}
}
