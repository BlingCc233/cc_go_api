package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"studyGo/global"
	"studyGo/models"
)

func CreatArticle(context *gin.Context) {
	var article = &models.Article{}

	if err := context.ShouldBindJSON(article); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	global.Db.AutoMigrate(article)
	global.Db.Create(article)

	var returnRes struct {
		Title   string `json:"title" binding:"required" gorm:"not null;unique"`
		Content string `json:"content" binding:"required"`
		Preview string `json:"preview" binding:"required"`
	}

	returnRes.Preview = article.Content[:1]
	returnRes.Title = article.Title
	returnRes.Content = article.Content

	context.JSON(http.StatusCreated, returnRes)
}

func GetArticles(context *gin.Context) {
	var articles []models.Article
	global.Db.Find(&articles)
	var Req struct {
		Page     uint `json:"page"`
		PageSize uint `json:"pageSize"`
	}
	err := context.ShouldBindJSON(&Req)
	if err != nil {
		Req.Page = 1
		Req.PageSize = uint(len(articles))
	}
	index := (Req.Page - 1) * Req.PageSize

	context.JSON(http.StatusOK, gin.H{"data": articles[index : index+Req.PageSize]})

}

func GetArticleById(context *gin.Context) {
	var article models.Article
	id := context.Param("id")
	if err := global.Db.First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": article})
}
