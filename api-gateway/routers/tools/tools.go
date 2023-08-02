package tools

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/models"
	"Api-Gateway-lcs42/utils"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func GetRequest(url, path string) (models.DtoResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.DtoResponse{}, errors.New("error to mount new request")
	}

	jwt, err := GetJwt()
	if err != nil {
		return models.DtoResponse{}, nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", jwt.Token)

	res, err := client.Do(req)
	if err != nil {
		return models.DtoResponse{}, nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.DtoResponse{}, nil
	}

	data, err := utils.ParseDtoResponse(body)
	if err != nil {
		return models.DtoResponse{}, nil
	}

	var buff []interface{}
	buff = append(buff, data)
	response := models.DtoResponse{
		Message:    "msg",
		Id:         uuid.NewString(),
		Data:       buff,
		StatusCode: 200,
	}

	return response, nil
}

func GetJwt(sevrer, authKey string) (*models.AuthJwt, error) {
	// TODO: aqui deve conter todas as url's que devolvem um jwt
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", config.SERVER1_AUTH_KEY)

	if err != nil {
		return &models.AuthJwt{}, errors.New("error to get jwt")
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return &models.AuthJwt{}, errors.New("error to get jwt")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return &models.AuthJwt{}, errors.New("error to get jwt")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &models.AuthJwt{}, errors.New("error to get jwt")
	}

	response, err := utils.ParseJwt(body)
	if err != nil {
		return &models.AuthJwt{}, errors.New("error to get jwt")
	}

	return response, nil
}
