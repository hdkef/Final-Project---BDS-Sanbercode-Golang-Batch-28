package swagmodel

type GetRating struct {
	ID        uint
	ArticleID uint
	Star      uint
}

type GetRatingSum struct {
	ID        uint
	ArticleID uint
	AvgStar   uint
}

type InputRating struct {
	Star uint
}
