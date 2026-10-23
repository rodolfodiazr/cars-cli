package commands

import (
	"context"
	"fmt"
	"os"
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
	fmt.Fprintln(w, "ID\tMAKE\tMODEL\tYEAR")

	for _, car := range cars {
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\n",
			car.ID, car.Make, car.Model, car.Year)
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

	fmt.Printf("ID:       %s\n", result.ID)
	fmt.Printf("Make:     %s\n", result.Make)
	fmt.Printf("Model:    %s\n", result.Model)
	fmt.Printf("Color:    %s\n", result.Color)
	fmt.Printf("Category: %s\n", result.Category)
	fmt.Printf("Year:     %d\n", result.Year)
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
