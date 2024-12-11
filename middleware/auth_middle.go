package middleware

import (
	"fmt"
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
		username, err := utils.ParseJWT(token)
		if err != nil {
			fmt.Printf("%v", err)
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			context.Abort()
			return
		}
		context.Set("username", username)
		context.Next()

	}
}
