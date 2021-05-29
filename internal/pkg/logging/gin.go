package logging

import (
	"github.com/gin-gonic/gin"
	"time"
)

func GinZapLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		path := ctx.Request.URL.Path
		ctx.Next()

		end := time.Now()
		duration := end.Sub(start)

		for _, e := range ctx.Errors.Errors() {
			gLogger.Error(e)
		}

		gLogger.Infow(path,
			"method", ctx.Request.Method,
			"status", ctx.Writer.Status(),
			"duration", duration.String(),
			"ip", ctx.ClientIP(),
		)
	}
}
