package utils

import (
	"encoding/json"
	"fmt"

	"github.com/api-server/lcs42/models"
)

// ParseToBytes try to execute a Marshal
func ParseToBytes(data interface{}) ([]byte, error) {
	dataBuffer, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}
	return dataBuffer, nil
}

func ParseToModels(data []byte) (models.User, error) {
	var user models.User
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return models.User{}, err
	}
	return user, nil

}
