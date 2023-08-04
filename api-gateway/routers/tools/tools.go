package tools

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/models"
	"Api-Gateway-lcs42/utils"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func GetRequest(url, path, token string) (models.DtoResponse, error) {
	client := &http.Client{}
	// * temp, the url should be complete at this point
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/all", url, path), nil)
	if err != nil {
		return models.DtoResponse{}, errors.New("error to mount new request")
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		return models.DtoResponse{}, nil
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return models.DtoResponse{}, fmt.Errorf("server do not respond 200")
	}

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
		Message:    "Sucess",
		Id:         uuid.NewString(),
		Data:       buff,
		StatusCode: 200,
	}

	return response, nil
}

func GetJwt(server string) (*models.AuthJwt, error) {
	var url string
	var authKey string

	// TODO: add url's .env
	switch server {
	case "server1":
		url = "http://127.0.0.1:2001/server1"
		authKey = config.SERVER1_AUTH_KEY
	case "server2":
		url = "http://127.0.0.1:8000/authentication"
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", authKey)
	if err != nil {
		return &models.AuthJwt{}, errors.New("error to mount jwt request")
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return &models.AuthJwt{}, errors.New("request error jwt")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return &models.AuthJwt{}, errors.New("server do not respond as expected")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &models.AuthJwt{}, errors.New("fail to read server response")
	}

	response, err := utils.ParseJwt(body)
	if err != nil {
		return &models.AuthJwt{}, errors.New("fail to parse jwt")
	}

	return response, nil
}
