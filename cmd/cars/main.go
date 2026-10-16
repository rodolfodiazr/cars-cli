package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/rodolfodiazr/cars-cli/internal/car"
	"github.com/rodolfodiazr/cars-cli/internal/client"
)

func main() {
	fmt.Println("CARS - COMMAND CLI")
	fmt.Println("------------------")
	fmt.Println("------------------")
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("------------------")
	fmt.Println("------------------")
}

func run() error {
	ctx := context.Background()

	sdk := client.NewCarClient(http.DefaultClient)

	// Retrieve the list of cars from the Cars API
	cars, err := sdk.ListCars(ctx, "", "", 0)
	if err != nil {
		return err
	}

	for index, c := range cars {
		fmt.Printf("%d | %s | %s | %s | %d\n", (index + 1), c.ID, c.Make, c.Model, c.Year)
	}

	// Retrieve a car from the Cars API
	c, err := sdk.GetCar(ctx, "1I3XJRLLC")
	if err != nil {
		return err
	}

	fmt.Printf("\n\n%s | %s | %s | %d\n\n", c.ID, c.Make, c.Model, c.Year)

	// Create a new car
	newCar := car.Car{
		Make:     "Ford",
		Model:    "F10",
		Color:    "Silver",
		Category: "Truck",
		Year:     2010,
	}
	carID, err := sdk.CreateCar(ctx, newCar)
	if err != nil {
		return err
	}

	// Update an existing car
	updatedCar := car.Car{
		Make:     "Ford",
		Model:    "F10",
		Color:    "Silver",
		Category: "Truck",
		Year:     2012,
	}
	if err := sdk.UpdateCar(ctx, carID, updatedCar); err != nil {
		return err
	}

	return nil
}
