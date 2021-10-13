package swagmodel

type GetDraft struct {
	ID        uint
	CreatorID uint
	Title     string
	ImageURL  string
	Body      string
}

type InputDraft struct {
	ID        uint
	CreatorID uint
	Title     string
	ImageURL  string
	Body      string
}
