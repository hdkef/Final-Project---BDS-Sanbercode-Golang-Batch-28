package swagmodel

type LoginPayload struct {
	Username string
	Password string
}

type LoginResponse struct {
	TokenType string
	Token     string
}

type ChPwd struct {
	Password string
}
