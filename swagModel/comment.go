package swagmodel

type GetComment struct {
	ID        uint   `json:"id" example:"1"`
	ArticleID uint   `json:"article_id" example:"1"`
	CreatorID uint   `json:"creator_id" example:"1"`
	Body      string `json:"body" example:"this is fun to read >_<"`
}

type InputComment struct {
	Body string `json:"body" example:"this is fun to read >_<"`
}
