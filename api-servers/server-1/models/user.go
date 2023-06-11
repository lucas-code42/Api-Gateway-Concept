package models

// User model
type User struct {
	Id    string `json:"id,omitempty" redis:"userId"`
	Name  string `json:"name" binding:"required" redis:"userName"`
	Email string `json:"email" binding:"required" redis:"userEmail"`
}


// TODO: CREATE METHODS TO VALIDADE ALL FIELDS