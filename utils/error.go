package utils

// JsonErr struct for error JSON response
type JsonErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
