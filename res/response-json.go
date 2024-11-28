package res

import "ujiTest/models"

type GetResponse struct {
	Status     int           `json:"status"`
	Count      int           `json:"count"`
	TotalCount int           `json:"totalCount"`
	Data       []models.Item `json:"data"`
}

type Status struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
