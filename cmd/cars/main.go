package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

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
	/*inputs := car.Car{
		Make:     "Ford",
		Model:    "--",
		Color:    "Red",
		Category: "--",
		Year:     2010,
	}
	if err := sdk.CreateCar(ctx, inputs); err != nil {
		return err
	}*/
	return nil
}
