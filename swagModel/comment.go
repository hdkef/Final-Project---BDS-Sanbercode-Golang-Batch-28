package swagmodel

type GetComment struct {
	ID        uint
	ArticleID uint
	CreatorID uint
	Body      string
}

type InputComment struct {
	Body string
}
