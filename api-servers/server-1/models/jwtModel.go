package models

type JwtModel struct {
	Token string `json:"token" binding:"required"`
}
