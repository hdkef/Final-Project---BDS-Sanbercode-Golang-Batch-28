package models

type Inbox struct {
	ID         uint
	CreatorID  uint
	ReceiverID uint
	Body       string
}
