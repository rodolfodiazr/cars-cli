package car

// Car represents a car in the application.
type Car struct {
	ID string

	Make     string
	Model    string
	Color    string
	Category string
	Year     int

	Package *string
	Mileage *int64
	Price   *int64
}

// Cars represents a list of car.
type Cars []Car

// CreateCar represents the information required to create a car.
type CreateCar struct {
	Make     string
	Model    string
	Color    string
	Category string
	Year     int
}

// UpdateCar represents the information used to update a car.
type UpdateCar struct {
	Make     string
	Model    string
	Color    string
	Category string
	Year     int
}
