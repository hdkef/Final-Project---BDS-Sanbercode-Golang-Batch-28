package models

type Article struct {
	ID        uint
	CreatorID uint
	Title     string
	ImageURL  string
	Body      string
}
