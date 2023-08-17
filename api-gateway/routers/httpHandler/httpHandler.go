package httpHandler

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

type DoRequest interface {
	Request(models.RequestHost) (models.DtoResponse, error)
	GetJwt(server string) (*models.AuthJwt, error)
}

func GetJwt(server string) (*models.AuthJwt, error) {
	var url string
	var authKey string

	switch server {
	case "server1":
		url = config.SERVER1_AUTH_PATH
		authKey = config.SERVER1_AUTH_KEY
	case "server2":
		url = config.SERVER2_AUTH_PATH
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", authKey)
	if err != nil {
		return &models.AuthJwt{}, errors.New("error to mount jwt request")
	}
	req.Header.Add("Authorization", authKey)

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

func Request(request models.RequestHost) (models.DtoResponse, error) {
	var p io.Reader
	if request.Payload == nil {
		p = nil
	} else {
		p = request.Payload
	}

	client := &http.Client{}
	req, err := http.NewRequest(
		request.Method,
		fmt.Sprintf("%s/%s", request.Url, request.Path),
		p,
	)

	if err != nil {
		return models.DtoResponse{}, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", request.Token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return models.DtoResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return models.DtoResponse{}, err
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
