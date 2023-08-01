package tools

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/models"
	"Api-Gateway-lcs42/utils"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func RequestServer2(method string, path string, payload *bytes.Buffer) (models.DtoResponse, error) {
	start := time.Now()

	jwt, err := GetAuthentication()
	if err != nil {
		return models.DtoResponse{}, err
	}

	url := fmt.Sprintf("%s/%s", config.DEFAULT_HOST_SERVER2, config.SERVER2_PATH)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return models.DtoResponse{}, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", jwt.Token)

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
		Message:       "msg",
		Id:            uuid.NewString(),
		Data:          buff,
		StatusCode:    200,
		ExecutionTime: time.Duration(time.Since(start).Milliseconds()),
	}

	return response, nil
}
