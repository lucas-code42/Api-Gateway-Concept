package models

import "time"

type DtoResponse struct {
	Message       string        `json:"message"`
	Id            string        `json:"id"`
	Data          []interface{} `json:"data"`
	StatusCode    int           `json:"statusCode"`
	ExecutionTime time.Duration `json:"miliSeconds"`
}

// TODO ajustar auth de uma forma que fique instanciado em apenas uma variável
type AuthJwt struct {
	Token string `json:"token"`
}
