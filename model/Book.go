package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func (b *Book) TableName() string {
	return "booki"
}
