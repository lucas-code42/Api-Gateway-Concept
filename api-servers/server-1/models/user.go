package models

import (
	"net/mail"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

// User model
type User struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

var (
	UPDATE_USER = "UPDATE"
	DELETE_USER = "DELETE"
	GET_USER    = "GET"
)

func NewUser() *User {
	return &User{}
}

func (u *User) ValidadeUserStruct(handler string) bool {
	switch handler {
	case UPDATE_USER:
		validUiid, err := uuid.Parse(u.Id)
		if err != nil {
			return false
		}
		u.Id = validUiid.String()

	case GET_USER:
		_, err := uuid.Parse(u.Id)
		if err != nil {
			return false
		} else {
			return true
		}
	}

	validAddr, err := mail.ParseAddress(u.Email)
	if err != nil {
		return false
	}
	u.Email = strings.TrimSpace(validAddr.Address)

	if onlyLetters(u.Name) {
		u.Name = strings.TrimSpace(u.Name)
	} else {
		return false
	}

	return true
}

func onlyLetters(data string) bool {
	for _, v := range data {
		if !unicode.IsLetter(v) {
			return false
		}
	}
	return true
}
