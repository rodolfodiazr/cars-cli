package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rodolfodiazr/cars-cli/internal/client"
	"github.com/rodolfodiazr/cars-cli/internal/commands"
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
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	apiClient := client.New(
		httpClient,
		"http://localhost:8080",
	)

	if err := commands.Create(ctx); err != nil {
		return err
	}

	cars, err := apiClient.ListCars(ctx)
	if err != nil {
		return err
	}
	fmt.Println("cars: ", cars)

	if err := commands.Get(ctx, "1"); err != nil {
		return err
	}
	if err := commands.Update(ctx, "1"); err != nil {
		return err
	}
	return commands.Delete(ctx, "1")
}
