package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"studyGo/global"
	"studyGo/models"
	"studyGo/utils"
)

func Register(context *gin.Context) {
	var user *models.User = &models.User{}

	if err := context.ShouldBindJSON(user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = hashedPassword

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	global.Db.AutoMigrate(*user)
	global.Db.Create(user)

	context.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(context *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	global.Db.Where("username = ?", input.Username).First(&user)
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, _ := utils.GenerateJWT(user.Username)
	context.JSON(http.StatusOK, gin.H{"token": token})
}

func TokenVerify(context *gin.Context) {
	token := context.GetHeader("Authorization")
	username, _ := utils.ParseJWT(token)
	// 从数据库中查询用户名是否存在
	var user models.User
	global.Db.AutoMigrate(user)
	global.Db.Where("username = ?", username).First(&user)
	if user.Username == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
}
