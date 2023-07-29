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
)

func GetAuthentication() (*models.AuthJwt, error) {
	url := fmt.Sprintf("%s/authenticate", config.DEFAULT_HOST_SERVER2)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
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
