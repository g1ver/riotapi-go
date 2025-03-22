package client

import (
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int
	Message    string
	Headers    http.Header
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error: status=%d, message=%s", e.StatusCode, e.Message)
}
