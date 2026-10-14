package client

type CarRequest struct {
	Make     string `json:"make"`
	Model    string `json:"model"`
	Color    string `json:"color"`
	Category string `json:"category"`
	Year     int    `json:"year"`
}
