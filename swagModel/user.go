package swagmodel

type GetUser struct {
	ID        uint
	Username  string
	Password  string
	Role      string
	Bio       string
	AvatarURL string
}

type GetUserPublic struct {
	ID        uint
	Username  string
	Role      string
	Bio       string
	AvatarURL string
}

type InputUser struct {
	Username string
	Password string
	Role     string
}
