package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rodolfodiazr/cars-cli/commands"
	"github.com/rodolfodiazr/cars-cli/internal/car"
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
		if len(os.Args) != 2 {
			return fmt.Errorf("usage: cars list")
		}
		return cars.List(ctx, "", "", 0)

	case "get":
		if len(os.Args) != 4 || os.Args[2] != "-i" {
			return fmt.Errorf("usage: cars get -i <carID>")
		}
		return cars.Get(ctx, os.Args[3])

	case "create":
		input, err := toCreateCar(os.Args[2:])
		if err != nil {
			return fmt.Errorf("usage: cars create -make <make> -model <model> -color <color> -category <category> -year <year>")
		}
		return cars.Create(ctx, input)

	case "update":
		input, err := toUpdateCar(os.Args[2:])
		if err != nil {
			return fmt.Errorf("usage: cars update -i <carID> -make <make> -model <model> -color <color> -category <category> -year <year>")
		}
		return cars.Update(ctx, os.Args[3], input)

	case "delete":
		if len(os.Args) != 4 || os.Args[2] != "-i" {
			return fmt.Errorf("usage: cars delete -i <carID>")
		}
		return cars.Delete(ctx, os.Args[3])

	default:
		return fmt.Errorf("unknown command %q: available commands: list, get, create, update, delete", os.Args[1])
	}
}

func toCreateCar(args []string) (car.CreateCar, error) {
	flags := flag.NewFlagSet("create", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	make := flags.String("make", "", "make")
	model := flags.String("model", "", "model")
	color := flags.String("color", "", "color")
	category := flags.String("category", "", "category")
	year := flags.Int("year", 0, "year")

	if err := flags.Parse(args); err != nil {
		return car.CreateCar{}, err
	}

	return car.CreateCar{
		Make:     *make,
		Model:    *model,
		Color:    *color,
		Category: *category,
		Year:     *year,
	}, nil
}

func toUpdateCar(args []string) (car.UpdateCar, error) {
	flags := flag.NewFlagSet("update", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	_ = flags.String("i", "", "id")
	make := flags.String("make", "", "make")
	model := flags.String("model", "", "model")
	color := flags.String("color", "", "color")
	category := flags.String("category", "", "category")
	year := flags.Int("year", 0, "year")

	if err := flags.Parse(args); err != nil {
		return car.UpdateCar{}, err
	}

	return car.UpdateCar{
		Make:     *make,
		Model:    *model,
		Color:    *color,
		Category: *category,
		Year:     *year,
	}, nil
}
