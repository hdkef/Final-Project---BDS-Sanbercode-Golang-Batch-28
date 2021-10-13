package models

type Draft struct {
	ID        uint
	CreatorID uint
	Title     string
	ImageURL  string
	Body      string
}
