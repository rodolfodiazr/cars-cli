package commands

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/rodolfodiazr/cars-cli/internal/car"
	"github.com/rodolfodiazr/cars-cli/internal/client"
)

// ListCarsParams represents the parameters for the list command.
type ListCarsParams struct {
	Make  string
	Model string
	Year  int
}

// UpdateCarParams represents the parameters for the update command.
type UpdateCarParams struct {
	ID       string
	Make     string
	Model    string
	Color    string
	Category string
	Year     int
	Package  *string
	Mileage  *int64
	Price    *int64
}

// CreateCarParams represents the parameters for the create command.
type CreateCarParams struct {
	Make     string
	Model    string
	Color    string
	Category string
	Year     int
	Package  *string
	Mileage  *int64
	Price    *int64
}

// Cars provides commands for managing cars.
type Cars struct {
	client client.Car
}

// NewCars creates a new Cars command handler using the provided car client.
func NewCars(client client.Car) Cars {
	return Cars{client: client}
}

// List retrieves and displays cars matching the provided filters.
func (c *Cars) List(ctx context.Context, params ListCarsParams) error {
	cars, err := c.client.ListCars(ctx, params.Make, params.Model, params.Year)
	if err != nil {
		return fmt.Errorf("list cars: %w", err)
	}

	if len(cars) == 0 {
		fmt.Println("No cars found.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tMAKE\tMODEL\tCOLOR\tCATEGORY\tYEAR\tPACKAGE\tMILEAGE\tPRICE")

	for _, car := range cars {
		pkg := "--"
		if car.Package != nil {
			pkg = *car.Package
		}
		mileage := "--"
		if car.Mileage != nil {
			mileage = formatInt64(*car.Mileage)
		}
		price := "--"
		if car.Price != nil {
			price = formatInt64(*car.Price)
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%d\t%s\t%s\t%s\n",
			car.ID, car.Make, car.Model, car.Color, car.Category, car.Year, pkg, mileage, price)
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("display cars: %w", err)
	}

	fmt.Printf("%d cars found.\n", len(cars))
	return nil
}

// Get retrieves and displays a car by its ID.
func (c *Cars) Get(ctx context.Context, carID string) error {
	result, err := c.client.GetCar(ctx, carID)
	if err != nil {
		return fmt.Errorf("get car %s: %w", carID, err)
	}

	pkg := "--"
	if result.Package != nil {
		pkg = *result.Package
	}
	mileage := "--"
	if result.Mileage != nil {
		mileage = formatInt64(*result.Mileage)
	}
	price := "--"
	if result.Price != nil {
		price = formatInt64(*result.Price)
	}

	fmt.Printf("ID:       %s\n", result.ID)
	fmt.Printf("Make:     %s\n", result.Make)
	fmt.Printf("Model:    %s\n", result.Model)
	fmt.Printf("Color:    %s\n", result.Color)
	fmt.Printf("Category: %s\n", result.Category)
	fmt.Printf("Year:     %d\n", result.Year)
	fmt.Printf("Package:  %s\n", pkg)
	fmt.Printf("Mileage:  %s\n", mileage)
	fmt.Printf("Price:    %s\n", price)
	return nil
}

// Create creates a car using the provided parameters.
func (c *Cars) Create(ctx context.Context, params CreateCarParams) error {
	req := car.CreateCar{
		Make:     params.Make,
		Model:    params.Model,
		Color:    params.Color,
		Category: params.Category,
		Year:     params.Year,
		Package:  params.Package,
		Mileage:  params.Mileage,
		Price:    params.Price,
	}

	carID, err := c.client.CreateCar(ctx, req)
	if err != nil {
		return fmt.Errorf("create car: %w", err)
	}

	return nil
}

// Update updates a car using the provided parameters.
func (c *Cars) Update(ctx context.Context, params UpdateCarParams) error {
	req := car.UpdateCar{
		Make:     params.Make,
		Model:    params.Model,
		Color:    params.Color,
		Category: params.Category,
		Year:     params.Year,
		Package:  params.Package,
		Mileage:  params.Mileage,
		Price:    params.Price,
	}

	if err := c.client.UpdateCar(ctx, params.ID, req); err != nil {
		return fmt.Errorf("update car %s: %w", params.ID, err)
	}

	fmt.Printf("Car updated successfully (ID: %s)", params.ID)

	return nil
}

// Delete deletes a car by its ID.
func (c *Cars) Delete(ctx context.Context, carID string) error {
	if err := c.client.DeleteCar(ctx, carID); err != nil {
		return fmt.Errorf("delete car %s: %w", carID, err)
	}

	fmt.Printf("Car deleted successfully (ID: %s)", carID)
	return nil
}

func formatInt64(value int64) string {
	s := strconv.FormatInt(value, 10)
	for i := len(s) - 3; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}
	return s
}
