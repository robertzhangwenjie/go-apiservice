package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"code":   200,
	})
}
