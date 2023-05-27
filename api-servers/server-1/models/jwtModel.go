package models

// JwtModel model
type JwtModel struct {
	Token string `json:"token" binding:"required"`
}
