package models

type Item struct {
	Code   string   `json:"code"`
	Name   string   `json:"name"`
	Model  string   `json:"model"`
	Tech   []string `json:"tech"`
	Status string   `json:"status"`
}
