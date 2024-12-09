package models

import "gorm.io/gorm"

// Config 表, for sql
type Config struct {
	gorm.Model

	Host     string `gorm:"unique;not null" json:"host"`
	Port     string `gorm:"unique;not null" json:"port"`
	User     string `gorm:"unique;not null" json:"user"`
	Password string `gorm:"unique;not null" json:"password"`
}
