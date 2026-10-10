package car

import (
	"fmt"
	"net/url"
)

type Car struct {
	ID    string
	Make  string
	Model string
	Year  int
}

type CarResponse struct {
	ID    string `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

type CarsResponse []CarResponse

// CarResponseError represents a generic HTTP error returned by the Cars API.
type CarResponseError struct {
	// StatusCode is the HTTP status code returned by the API.
	StatusCode int

	// RequestURL is the full URL that was called.
	RequestURL *url.URL

	// ResponseBody contains the raw response body returned by the API.
	ResponseBody string
}

// Error implements the error interface.
func (e CarResponseError) Error() string {
	return fmt.Sprintf("Status %d: Request URL: %v Error Body: %s", e.StatusCode, e.RequestURL, e.ResponseBody)
}
