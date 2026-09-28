package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rodolfodiazr/cars-cli/internal/commands"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if err := commands.List(context.Background()); err != nil {
		return err
	}
	if err := commands.Get(context.Background(), "1"); err != nil {
		return err
	}
	if err := commands.Create(context.Background()); err != nil {
		return err
	}
	return commands.Update(context.Background(), "1")
}
