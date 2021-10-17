package swagmodel

type GetMedia struct {
	ID        uint
	CreatorID uint
	Caption   string
	Alt       string
	Url       string
}

type InputMedia struct {
	Caption string `json:"caption" form:"caption" format:"string"`
	Alt     string `json:"alt" form:"alt" format:"string"`
	File    string `json:"file" form:"file" format:"binary"`
}
