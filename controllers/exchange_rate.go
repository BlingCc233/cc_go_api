package controllers

import (
	"errors"
	"net/http"
	"studyGo/global"
	"studyGo/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateExchangeRate(context *gin.Context) {
	var exchangeRate models.ExchangeRate

	if err := context.ShouldBindJSON(&exchangeRate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exchangeRate.Date = time.Now()

	if err := global.Db.AutoMigrate(&exchangeRate); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.Create(&exchangeRate).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, exchangeRate)
}

func GetExchangeRates(context *gin.Context) {
	var exchangeRates []models.ExchangeRate

	if err := global.Db.Find(&exchangeRates).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	context.JSON(http.StatusOK, exchangeRates)
}
