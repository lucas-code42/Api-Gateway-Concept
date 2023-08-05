package models

import "time"

type DtoResponse struct {
	Message       string        `json:"message"`
	Id            string        `json:"id"`
	Data          []interface{} `json:"data"`
	StatusCode    int           `json:"statusCode"`
	ExecutionTime time.Duration `json:"miliSeconds"`
}

type AuthJwt struct {
	Token string `json:"token"`
}
