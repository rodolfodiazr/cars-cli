package client

// CarResponse represents a car as returned by the API.
type CarResponse struct {
	ID    string `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

// CarsResponse represents a collection of cars returned by the API.
type CarsResponse []CarResponse

// CarValidationError ...
type CarValidationError struct {
	Error   string `json:"error"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}
