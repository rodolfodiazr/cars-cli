package client

// CreateCarRequest represents the JSON payload sent to create a car.
type CreateCarRequest struct {
	Make     string  `json:"make"`
	Model    string  `json:"model"`
	Color    string  `json:"color"`
	Category string  `json:"category"`
	Year     int     `json:"year"`
	Package  *string `json:"package"`
	Mileage  *int64  `json:"mileage"`
	Price    *int64  `json:"price"`
}

// UpdateCarRequest represents the JSON payload sent to update a car.
type UpdateCarRequest struct {
	Make     string `json:"make"`
	Model    string `json:"model"`
	Color    string `json:"color"`
	Category string `json:"category"`
	Year     int    `json:"year"`
}
