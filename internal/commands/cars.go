package commands

import (
	"context"
)

func List(ctx context.Context) error {
	// call client.ListCars()
	// display results
	return nil
}

func Get(ctx context.Context, carID string) error {
	// call client.GetCar()
	// display result
	return nil
}

func Create(ctx context.Context) error {
	// call client.CreateCar()
	return nil
}

func Update(ctx context.Context, carID string) error {
	// call client.UpdateCar()
	return nil
}

func Delete(ctx context.Context, carID string) error {
	// call client.DeleteCar()
	return nil
}
