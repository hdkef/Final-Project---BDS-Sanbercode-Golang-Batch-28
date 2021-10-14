package models

type Comment struct {
	ID        uint   `json:"id"`
	ArticleID uint   `json:"article_id"`
	CreatorID uint   `json:"creator_id"`
	Body      string `json:"body_id"`
}
