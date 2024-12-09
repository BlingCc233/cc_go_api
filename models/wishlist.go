package models

type WishList struct {
	ID      uint   `gorm:"primarykey" json:"id"`
	Content string `json:"content" binding:"required"`
	Checked bool   `json:"checked"`
	Creater string `json:"creater" binding:"required"`
}
