package swagmodel

type LoginPayload struct {
	Username string `json:"username" example:"adm"`
	Password string `json:"password" example:"superhardpassword"`
}

type LoginResponse struct {
	TokenType string `json:"id" example:"new"`
	Token     string `json:"token" example:"rANDomStringJwtTOKenGenERatedWIthSecRET"`
}

type Register struct {
	Username  string `json:"username" example:"adm"`
	Password  string `json:"password" example:"superhardpassword"`
	Bio       string `json:"bio" example:"i am admin"`
	AvatarURL string `json:"avatar_url" example:"https://pbs.twimg.com/profile_images/1363210545118150659/Uo-XiGtv_400x400.jpg"`
}

type ChPwd struct {
	Password string `json:"password" example:"anothersuperhardpassword"`
}
