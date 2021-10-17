package swagmodel

type GetDraft struct {
	ID        uint   `json:"id" example:"1"`
	CreatorID uint   `json:"creator_id" example:"1"`
	Title     string `json:"title" example:"Review John Wick"`
	ImageURL  string `json:"image_url" example:"https://pbs.twimg.com/profile_images/1363210545118150659/Uo-XiGtv_400x400.jpg"`
	Body      string `json:"body" example:"<h1>REVIEW JOHN WICK</h1>Hello guys<br>...dst"`
}

type InputDraft struct {
	Title    string `json:"title" example:"Review John Wick"`
	ImageURL string `json:"image_url" example:"https://pbs.twimg.com/profile_images/1363210545118150659/Uo-XiGtv_400x400.jpg"`
	Body     string `json:"body" example:"<h1>REVIEW JOHN WICK</h1>Hello guys<br>...dst"`
}
