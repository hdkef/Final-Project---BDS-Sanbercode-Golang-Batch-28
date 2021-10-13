package models

type User struct {
	ID        uint
	Username  string
	Password  string
	Role      string
	Bio       string
	AvatarURL string
}

func (m *User) GenerateToken() (string, error) {
	//get user detail from db filter by username
	//compare user password from db with password from payload
	//generate jwt token
	return "", nil
}
