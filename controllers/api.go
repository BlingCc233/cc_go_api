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
	var count int64
	global.Db.Model(&qcs).Count(&count)

	rand.Seed(time.Now().UnixNano())
	id := rand.Int63n(count) + 1

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

// 获取谁是卧底词
func GetSswd(context *gin.Context) {
	var wd models.WhosSpy
	var count int64
	global.Db.Model(&wd).Count(&count)
	rand.Seed(time.Now().UnixNano())
	id := rand.Int63n(count) + 1
	err := global.Db.Where("id = ?", id).First(&wd).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get word"})
		return
	}

	// 创建一个不包含 ID 的新对象
	responseData := gin.H{
		"word1": wd.Word1,
		"word2": wd.Word2,
	}

	// 返回去掉 ID 的数据
	context.JSON(http.StatusOK, gin.H{"data": responseData})
}
