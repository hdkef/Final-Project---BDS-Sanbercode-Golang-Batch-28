package models

type Comment struct {
	ID        uint
	ArticleID uint
	CreatorID uint
	Body      string
}
