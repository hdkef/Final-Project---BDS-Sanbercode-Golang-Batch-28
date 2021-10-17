package swagmodel

type Response struct {
	Success bool   `json:"success" example:"true"`
	Msg     string `json:"msg" example:"article posted"`
}
