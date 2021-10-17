package swagmodel

type GetUser struct {
	ID        uint   `json:"id" example:"1"`
	Username  string `json:"username" example:"admin"`
	Password  string `json:"password" example:"superhardpassword"`
	Role      string `json:"role" example:"regular"`
	Bio       string `json:"bio" example:"i am admin"`
	AvatarURL string `json:"avatar_url" example:"https://pbs.twimg.com/profile_images/1363210545118150659/Uo-XiGtv_400x400.jpg"`
}

type GetUserPublic struct {
	ID        uint   `json:"id" example:"1"`
	Username  string `json:"username" example:"admin"`
	Role      string `json:"role" example:"regular"`
	Bio       string `json:"bio" example:"i am admin"`
	AvatarURL string `json:"avatar_url" example:"https://pbs.twimg.com/profile_images/1363210545118150659/Uo-XiGtv_400x400.jpg"`
}

type InputUser struct {
	Username  string `json:"username" example:"admin"`
	Password  string `json:"password" example:"superhardpassword"`
	Role      string `json:"role" example:"regular"`
	Bio       string `json:"bio" example:"i am admin"`
	AvatarURL string `json:"avatar_url" example:"https://pbs.twimg.com/profile_images/1363210545118150659/Uo-XiGtv_400x400.jpg"`
}
