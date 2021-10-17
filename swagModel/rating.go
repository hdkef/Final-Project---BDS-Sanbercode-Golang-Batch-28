package swagmodel

type GetRating struct {
	ID        uint `json:"id" example:"1"`
	ArticleID uint `json:"article_id" example:"1"`
	Star      uint `json:"sta" example:"5"`
}

type GetRatingSum struct {
	ArticleID uint `json:"article_id" example:"1"`
	AvgStar   uint `json:"avgstar" example:"5"`
}

type InputRating struct {
	Star uint `json:"star" example:"5"`
}
