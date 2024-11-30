package models

type Item struct {
	Code   string   `json:"code" validate:"required,min=10"`
	Name   string   `json:"name" validate:"required,email"`
	Model  string   `json:"model" validate:"required"`
	Tech   []string `json:"tech" validate:"required"`
	Status string   `json:"status" validate:"required"`
}

type ErrorInputResponse struct {
	FieldName string `json:"fieldName"`
	Message   string `json:"message"`
}
