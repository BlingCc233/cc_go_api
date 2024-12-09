package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"studyGo/utils"
)

func AuthMiddle() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			context.Abort()
			return
		}
		username := utils.ParseJWT(token)
		context.Set("username", username)
		context.Next()

	}
}
