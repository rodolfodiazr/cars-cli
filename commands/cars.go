package commands

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/rodolfodiazr/cars-cli/internal/car"
)

type CarClient interface {
	ListCars(ctx context.Context, make, model string, year int) (car.Cars, error)
}

type Cars struct {
	client CarClient
}

func NewCars(client CarClient) *Cars {
	return &Cars{client: client}
}

func (c *Cars) List(ctx context.Context, make, model string, year int) error {
	cars, err := c.client.ListCars(ctx, make, model, year)
	if err != nil {
		return fmt.Errorf("list cars: %w", err)
	}

	if len(cars) == 0 {
		fmt.Println("No cars found.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tMAKE\tMODEL\tYEAR")

	for _, car := range cars {
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\n",
			car.ID, car.Make, car.Model, car.Year)
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("display cars: %w", err)
	}

	fmt.Printf("%d cars found.\n", len(cars))
	return nil
}
