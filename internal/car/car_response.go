package car

// CarResponse ...
type CarResponse struct {
	ID       string
	Make     string
	Model    string
	Color    string
	Category string
	Year     int
}

// CarsResponse ...
type CarsResponse []CarResponse
