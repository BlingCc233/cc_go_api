package controllers

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"studyGo/global"
	"studyGo/models"
	"time"
)

func QCSLottery(context *gin.Context) {
	var qcs *models.QCS = &models.QCS{}

	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 随机生成一个1到100的ID
	id := rand.Intn(100) + 1

	err := global.Db.Where("id = ?", id).First(&qcs).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get QCS Lottery"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": qcs})
}
