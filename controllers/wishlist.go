package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studyGo/global"
	"studyGo/models"
)

func NewWish(context *gin.Context) {
	var wishlist = &models.WishList{}

	if err := context.ShouldBindJSON(wishlist); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username, exists := context.Get("username")
	if !exists {
		context.JSON(http.StatusBadRequest, gin.H{"error": "username not found in context"})
		return
	}

	global.Db.AutoMigrate(wishlist)
	global.Db.Create(wishlist)

	var returnRes struct {
		ID      uint   `json:"id" gorm:"not null;unique"`
		Content string `json:"content"`
		Checked bool   `json:"checked"`
		Creater string `json:"creater"`
	}

	wishlist.Creater = username.(string)
	returnRes.ID = wishlist.ID
	returnRes.Content = wishlist.Content
	returnRes.Checked = wishlist.Checked
	returnRes.Creater = wishlist.Creater

	context.JSON(http.StatusOK, returnRes)
}

func GetWish(context *gin.Context) {
	var wishlists []models.WishList
	global.Db.Find(&wishlists)
	context.JSON(http.StatusOK, gin.H{"data": wishlists})
}

func DeleteWish(context *gin.Context) {
	wishlist := models.WishList{}

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	wishlist.ID = uint(id)
	if err := global.Db.Delete(&wishlist).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Wish deleted successfully"})
}

func WishCheck(context *gin.Context) {
	wishlist := models.WishList{}
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	wishlist.ID = uint(id)
	var form struct {
		Checked bool `json:"checked"`
	}
	if err := context.ShouldBindJSON(&form); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := global.Db.Model(&wishlist).Update("checked", form.Checked).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Wish checked successfully"})
}
