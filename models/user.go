package models

type User struct {
	ID        uint
	Username  string
	Password  string
	Role      string
	Bio       string
	AvatarURL string
}
