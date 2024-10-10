package model

type HttpResponse[T any] struct {
	Success bool   `json:"success"`
	Data    *T     `json:"data"`
	Error   string `json:"error"`
}
