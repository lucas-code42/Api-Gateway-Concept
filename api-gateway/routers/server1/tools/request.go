package tools

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/models"
	"Api-Gateway-lcs42/utils"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func RequestServer1(method string, path string, payload io.Reader) (models.DtoResponse, error) {
	start := time.Now()

	url := fmt.Sprintf("%s/%s", config.DEFAULT_HOST_SERVER1, path)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
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
		Message:       "msg",
		Id:            uuid.NewString(),
		Data:          buff,
		StatusCode:    200,
		ExecutionTime: time.Duration(time.Since(start).Milliseconds()),
	}

	return response, nil
}
