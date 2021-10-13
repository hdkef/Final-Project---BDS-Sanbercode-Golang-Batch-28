package swagmodel

type GetInbox struct {
	ID         uint
	CreatorID  uint
	ReceiverID uint
	Body       string
}

type InputInbox struct {
	Body string
}
