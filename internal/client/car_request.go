package client

// CreateCarRequest represents the JSON payload sent to create a car.
type CreateCarRequest struct {
	Make     string `json:"make"`
	Model    string `json:"model"`
	Color    string `json:"color"`
	Category string `json:"category"`
	Year     int    `json:"year"`
}

// UpdateCarRequest represents the JSON payload sent to update a car.
type UpdateCarRequest struct {
	Make     string `json:"make"`
	Model    string `json:"model"`
	Color    string `json:"color"`
	Category string `json:"category"`
	Year     int    `json:"year"`
}
