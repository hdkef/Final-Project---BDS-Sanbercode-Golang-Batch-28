package models

type Response struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}
