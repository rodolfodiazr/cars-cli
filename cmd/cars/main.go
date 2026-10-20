package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/rodolfodiazr/cars-cli/commands"
	"github.com/rodolfodiazr/cars-cli/internal/client"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("missing command: usage: cars <command>")
	}

	ctx := context.Background()

	sdk := client.NewCarClient(
		http.DefaultClient,
		"http://localhost:8080",
	)

	cars := commands.NewCars(sdk)

	switch os.Args[1] {
	case "list":
		return cars.List(ctx, "", "", 0)
	default:
		return fmt.Errorf("unknown command %q: available command: list", os.Args[1])
	}
	/*
		// Create a new car
		newCar := car.CreateCar{
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
		fmt.Printf("\nCar created: %s | %s | %s | %d\n\n", carID, newCar.Make, newCar.Model, newCar.Year)

		// Update an existing car
		updatedCar := car.UpdateCar{
			Make:     "Ford",
			Model:    "F10",
			Color:    "Silver",
			Category: "Truck",
			Year:     2012,
		}
		if err := sdk.UpdateCar(ctx, carID, updatedCar); err != nil {
			return err
		}
		fmt.Printf("Car updated: %s | %s | %s | %d\n\n", carID, updatedCar.Make, updatedCar.Model, updatedCar.Year)

		// Retrieve a car from the Cars API
		c, err := sdk.GetCar(ctx, carID)
		if err != nil {
			return err
		}
		fmt.Printf("Car retrieved: %s | %s | %s | %d\n\n", c.ID, c.Make, c.Model, c.Year)

		// Delete a car from the Cars API
		if err := sdk.DeleteCar(ctx, carID); err != nil {
			return err
		}
		fmt.Printf("Car deleted: %s | %s | %s | %d\n\n", c.ID, c.Make, c.Model, c.Year)

		fmt.Println("Cars retrieved:")

		// Retrieve the list of cars from the Cars API
		cars, err := sdk.ListCars(ctx, "", "", 0)
		if err != nil {
			return err
		}

		for index, c := range cars {
			fmt.Printf("%d | %s | %s | %s | %d\n", (index + 1), c.ID, c.Make, c.Model, c.Year)
		}
	*/
}
