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
		if len(os.Args) != 2 {
			return fmt.Errorf("usage: cars list")
		}
		return cars.List(ctx, "", "", 0)

	case "get":
		if len(os.Args) != 4 || os.Args[2] != "-i" {
			return fmt.Errorf("usage: cars get -i <carID>")
		}
		return cars.Get(ctx, os.Args[3])

	default:
		return fmt.Errorf("unknown command %q: available commands: list, get", os.Args[1])
	}
}
