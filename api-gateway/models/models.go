package models

import (
	"bytes"
	"time"
)

// const para definir métodos http
const (
	GET    = 1
	POST   = 2
	PUT    = 3
	DELETE = 4
)

// DtoResponse dto padrão de saída da api
type DtoResponse struct {
	Message       string        `json:"message"`
	Id            string        `json:"id"`
	Data          []interface{} `json:"data"`
	StatusCode    int           `json:"statusCode"`
	ExecutionTime time.Duration `json:"miliSeconds"`
}

// AuthJwt
type AuthJwt struct {
	Token string `json:"token"`
}

// RequestHost conjunto de dados para fazer uma request a um servidor
type RequestHost struct {
	Url     string
	Path    string
	Token   string
	Method  string
	Payload *bytes.Buffer
}
