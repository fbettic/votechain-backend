package dto

type ErrorMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
