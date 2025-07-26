package models

type SuccessResponse struct {
	Success bool `json:"success"`
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
