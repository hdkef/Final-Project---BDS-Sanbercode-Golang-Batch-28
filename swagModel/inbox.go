package swagmodel

type GetInbox struct {
	ID         uint   `json:"id" example:"1"`
	CreatorID  uint   `json:"creator_id" example:"1"`
	ReceiverID uint   `json:"receiver_id" example:"2"`
	Body       string `json:"body" example:"nice article m8!"`
}

type InputInbox struct {
	Body string `json:"body" example:"nice article m8!"`
}
