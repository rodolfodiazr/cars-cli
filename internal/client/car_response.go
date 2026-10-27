package client

// CarResponse represents a car returned by the Cars API.
type CarResponse struct {
	CarID string `json:"id"`

	Make     string `json:"make"`
	Model    string `json:"model"`
	Color    string `json:"color"`
	Category string `json:"category"`
	Year     int    `json:"year"`

	Package *string `json:"package,omitempty"`
	Mileage *int64  `json:"mileage,omitempty"`
	Price   *int64  `json:"price,omitempty"`
}

// CarsResponse represents a collection of cars returned by the Cars API.
type CarsResponse []CarResponse

// CarValidationError represents a validation error returned by the Cars API.
type CarValidationError struct {
	Error   string `json:"error"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}
