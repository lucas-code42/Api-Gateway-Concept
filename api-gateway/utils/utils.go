package utils

import (
	"Api-Gateway-lcs42/models"
	"encoding/json"
)

func ParseDtoResponse(data []byte) (interface{}, error) {
	var buffer interface{}
	if err := json.Unmarshal(data, &buffer); err != nil {
		return nil, err
	}
	return buffer, nil
}

func ParseJwt(token []byte) (*models.AuthJwt, error) {
	var buffer models.AuthJwt
	if err := json.Unmarshal(token, &buffer); err != nil {
		return &models.AuthJwt{}, err
	}
	return &buffer, nil
}
