package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 设置允许的源
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的头部
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// 设置允许的方法
		context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if context.Request.Method == "OPTIONS" {
			fmt.Printf("OPTION")
			context.AbortWithStatus(http.StatusOK)
			return
		}

		context.Next()
	}
}
