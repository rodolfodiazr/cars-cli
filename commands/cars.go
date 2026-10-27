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

type Cars struct {
	client client.Car
}

func NewCars(client client.Car) Cars {
	return Cars{client: client}
}

func (c *Cars) List(ctx context.Context, make, model string, year int) error {
	cars, err := c.client.ListCars(ctx, make, model, year)
	if err != nil {
		return fmt.Errorf("list cars: %w", err)
	}

	if len(cars) == 0 {
		fmt.Println("No cars found.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tMAKE\tMODEL\tYEAR\tPACKAGE\tMILEAGE\tPRICE")

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

		fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%s\t%s\t%s\n",
			car.ID, car.Make, car.Model, car.Year, pkg, mileage, price)
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("display cars: %w", err)
	}

	fmt.Printf("%d cars found.\n", len(cars))
	return nil
}

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
	fmt.Printf("Package:  %s\n", pkg)
	fmt.Printf("Mileage:  %s\n", mileage)
	fmt.Printf("Price:    %s\n", price)
	return nil
}

func (c *Cars) Create(ctx context.Context, input car.CreateCar) error {
	carID, err := c.client.CreateCar(ctx, input)
	if err != nil {
		return fmt.Errorf("create car: %w", err)
	}

	fmt.Printf("Car created successfully (ID: %s): %s %s, %d\n",
		carID, input.Make, input.Model, input.Year)

	return nil
}

func (c *Cars) Update(ctx context.Context, carID string, input car.UpdateCar) error {
	if err := c.client.UpdateCar(ctx, carID, input); err != nil {
		return fmt.Errorf("update car %s: %w", carID, err)
	}

	fmt.Printf("Car %s updated successfully: %s %s, %d\n",
		carID, input.Make, input.Model, input.Year)

	return nil
}

func (c *Cars) Delete(ctx context.Context, carID string) error {
	if err := c.client.DeleteCar(ctx, carID); err != nil {
		return fmt.Errorf("delete car %s: %w", carID, err)
	}

	fmt.Printf("Car %s deleted successfully.\n", carID)
	return nil
}

func formatInt64(value int64) string {
	s := strconv.FormatInt(value, 10)

	for i := len(s) - 3; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}

	return s
}
