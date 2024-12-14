package controllers

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"studyGo/global"
	"studyGo/models"
	"studyGo/utils"
	"time"
)

func QCSLottery(context *gin.Context) {
	isJson := context.Query("return")

	var qcs models.QCS

	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 随机生成一个1到100的ID
	id := rand.Intn(100) + 1

	err := global.Db.Where("id = ?", id).First(&qcs).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get QCS Lottery"})
		return
	}

	if isJson == "json" {
		context.JSON(http.StatusOK, gin.H{"data": qcs})
	} else {
		// Generate image
		imageData, err := utils.GenerateQCSImage(qcs)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		context.Data(http.StatusOK, "image/jpeg", imageData)
	}
}

func XBImg(context *gin.Context) {
	content := context.Query("content")
	imageData, err := utils.GenerateXBImage(content)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.Data(http.StatusOK, "image/jpeg", imageData)
}

func YesNo(context *gin.Context) {
	ans := rand.Intn(2) == 1
	var res []string
	if ans {
		res = append(res, "1")
		res = append(res, "Yes")
		res = append(res, "是")
	} else {
		res = append(res, "0")
		res = append(res, "No")
		res = append(res, "否")
	}

	context.JSON(http.StatusOK, gin.H{"data": res})
}

func GetPhiB19(context *gin.Context) {
	session := context.Query("session")

	phigrosData, err := utils.GenPhiB19(session)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.Data(http.StatusOK, "image/jpeg", phigrosData)
}

func Text2QR(context *gin.Context) {
	text := context.Query("text")
	size := context.Query("size")
	intSize, err := strconv.Atoi(size)
	if err != nil || size == "" {
		size = "512"
		intSize, _ = strconv.Atoi(size)
	}
	imageData, err := utils.QRCode(text, intSize)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.Data(http.StatusOK, "image/png", imageData)
}
