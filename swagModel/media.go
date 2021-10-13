package swagmodel

type GetMedia struct {
	ID        uint
	CreatorID uint
	Caption   string
	Alt       string
	Url       string
}

type InputMedia struct {
	Caption string
	Alt     string
	Url     string
}
