package swagmodel

type GetRating struct {
	ID        uint
	ArticleID uint
	Star      uint
}

type InputRating struct {
	Star uint
}
