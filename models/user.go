package models

import "gorm.io/gorm"

// User 表, for sql
type User struct {
	gorm.Model

	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}
