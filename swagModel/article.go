package swagmodel

type GetArticle struct {
	ID        uint
	CreatorID uint
	Title     string
	ImageURL  string
	Body      string
}

type InputArticle struct {
	Title    string
	ImageURL string
	Body     string
}
