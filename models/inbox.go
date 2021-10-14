package models

type Inbox struct {
	ID         uint   `json:"id"`
	CreatorID  uint   `json:"creator_id"`
	ReceiverID uint   `json:"receiver_id"`
	Body       string `json:"body"`
}
