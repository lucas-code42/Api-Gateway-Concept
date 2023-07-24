package utils

import (
	"Api-Gateway-lcs42/models"
	"encoding/json"
)

func ParseDtoResponse(data []byte) (interface{}, error) {
	var buffer interface{}
	if err := json.Unmarshal(data, &buffer); err != nil {
		return models.DtoResponse{}, err
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

// TODO: improve the default response
// think of a way how to use dto model
// func MountResponse(message string, data []interface{}, stausCode int) map[string]any {
// 	return gin.H{
// 		"data": models.DtoResponse{
// 			Message:    message,
// 			Data:       data,
// 			Id:         uuid.NewString(),
// 			StatusCode: stausCode,
// 			Time:       time.Now().Format("2017.09.07 17:06:06"),
// 		},
// 	}
// }
