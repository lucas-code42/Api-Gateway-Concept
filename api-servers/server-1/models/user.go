package models

// User model
type User struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
