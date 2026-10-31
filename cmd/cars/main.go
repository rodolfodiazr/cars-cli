package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rodolfodiazr/cars-cli/commands"
	"github.com/rodolfodiazr/cars-cli/internal/client"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("missing command: usage: cars <command>")
	}

	ctx := context.Background()

	sdk := client.NewCarClient(
		http.DefaultClient,
		"http://localhost:8080",
	)

	cars := commands.NewCars(sdk)

	switch args[0] {
	case "list":
		params, err := parseListCarsArgs(args[1:])
		if err != nil {
			return fmt.Errorf("invalid list arguments: %w", err)
		}
		return cars.List(ctx, params)

	case "get":
		carID, err := parseCarIDArg(args[1:])
		if err != nil {
			return fmt.Errorf("invalid get arguments: %w", err)
		}
		return cars.Get(ctx, carID)

	case "create":
		params, err := parseCreateCarArgs(args[1:])
		if err != nil {
			return fmt.Errorf("invalid create arguments: %w", err)
		}
		return cars.Create(ctx, params)

	case "update":
		params, err := parseUpdateCarArgs(args[1:])
		if err != nil {
			return fmt.Errorf("invalid update arguments: %w", err)
		}
		return cars.Update(ctx, params)

	case "delete":
		carID, err := parseCarIDArg(args[1:])
		if err != nil {
			return fmt.Errorf("invalid delete arguments: %w", err)
		}
		return cars.Delete(ctx, carID)

	default:
		return fmt.Errorf("unknown command %q: available commands: list, get, create, update, delete", args[0])
	}
}

func parseCarIDArg(args []string) (string, error) {
	flags := flag.NewFlagSet("car", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	id := flags.String("id", "", "id")

	if err := flags.Parse(args); err != nil {
		return "", err
	}

	if flags.NArg() > 0 {
		return "", fmt.Errorf("unexpected arguments: %v", flags.Args())
	}

	if *id == "" {
		return "", fmt.Errorf("missing required flag: -id")
	}

	return *id, nil
}

func parseListCarsArgs(args []string) (commands.ListCarsParams, error) {
	flags := flag.NewFlagSet("list", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	carMake := flags.String("make", "", "make")
	carModel := flags.String("model", "", "model")
	carYear := flags.Int("year", 0, "year")

	if err := flags.Parse(args); err != nil {
		return commands.ListCarsParams{}, err
	}

	if flags.NArg() > 0 {
		return commands.ListCarsParams{}, fmt.Errorf("unexpected arguments: %v", flags.Args())
	}

	return commands.ListCarsParams{
		Make:  *carMake,
		Model: *carModel,
		Year:  *carYear,
	}, nil
}

func parseCreateCarArgs(args []string) (commands.CreateCarParams, error) {
	flags := flag.NewFlagSet("create", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	carMake := flags.String("make", "", "make")
	carModel := flags.String("model", "", "model")
	carColor := flags.String("color", "", "color")
	carCategory := flags.String("category", "", "category")
	carYear := flags.Int("year", 0, "year")

	if err := flags.Parse(args); err != nil {
		return commands.CreateCarParams{}, err
	}

	if flags.NArg() > 0 {
		return commands.CreateCarParams{}, fmt.Errorf("unexpected arguments: %v", flags.Args())
	}

	return commands.CreateCarParams{
		Make:     *carMake,
		Model:    *carModel,
		Color:    *carColor,
		Category: *carCategory,
		Year:     *carYear,
	}, nil
}

func parseUpdateCarArgs(args []string) (commands.UpdateCarParams, error) {
	flags := flag.NewFlagSet("update", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	carID := flags.String("id", "", "id")
	carMake := flags.String("make", "", "make")
	carModel := flags.String("model", "", "model")
	carColor := flags.String("color", "", "color")
	carCategory := flags.String("category", "", "category")
	carYear := flags.Int("year", 0, "year")

	if err := flags.Parse(args); err != nil {
		return commands.UpdateCarParams{}, err
	}

	if flags.NArg() > 0 {
		return commands.UpdateCarParams{}, fmt.Errorf("unexpected arguments: %v", flags.Args())
	}

	if *carID == "" {
		return commands.UpdateCarParams{}, fmt.Errorf("missing required flag: -id")
	}

	return commands.UpdateCarParams{
		ID:       *carID,
		Make:     *carMake,
		Model:    *carModel,
		Color:    *carColor,
		Category: *carCategory,
		Year:     *carYear,
	}, nil
}
