package models

import "gorm.io/gorm"

type Rating struct {
	ID        uint `json:"id"`
	ArticleID uint `json:"article_id"`
	Star      uint `json:"star"`
	gorm.Model
}
