package models

type Rating struct {
	ID        uint `json:"id"`
	ArticleID uint `json:"article_id"`
	Star      uint `json:"star"`
}
