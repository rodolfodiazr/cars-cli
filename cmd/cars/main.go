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
	carListResponse, err := sdk.ListCars(ctx)
	if err != nil {
		return err
	}

	// Retrieve the list of cars from the Cars API
	for index, c := range carListResponse {
		fmt.Printf("%d | %s | %s | %s | %d\n", (index + 1), c.ID, c.Make, c.Model, c.Year)
	}
	return nil
}
