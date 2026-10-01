package main

import (
	"context"
	"fmt"
	"os"

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
	if err := commands.Create(ctx); err != nil {
		return err
	}
	if err := commands.List(ctx); err != nil {
		return err
	}
	if err := commands.Get(ctx, "1"); err != nil {
		return err
	}
	if err := commands.Update(ctx, "1"); err != nil {
		return err
	}
	return commands.Delete(ctx, "1")
}
