package models

import "gorm.io/gorm"

type Media struct {
	ID        uint   `json:"id"`
	CreatorID uint   `json:"creator_id"`
	Caption   string `json:"caption"`
	Alt       string `json:"alt"`
	Url       string `json:"url"`
	gorm.Model
}
