package models

type DtoResponse struct {
	Message       string
	Time          string
	Id            string
	Data          []interface{}
	StatusCode    int
	ExecutionTime float64
}
