package commands

import (
	"context"
	"fmt"
)

func List(ctx context.Context) error {
	fmt.Printf("%s\t%s\t%s\t%s\n",
		"1",
		"Toyota",
		"Corolla",
		"2024",
	)
	return nil
}

func Get(ctx context.Context, carID string) error {
	fmt.Printf("%s\t%s\t%s\t%s\n",
		"1",
		"Toyota",
		"Corolla",
		"2024",
	)
	return nil
}

func Create(ctx context.Context) error {
	fmt.Println("Car created")
	return nil
}

func Update(ctx context.Context, carID string) error {
	fmt.Println("Car updated")
	return nil
}

func Delete(ctx context.Context, carID string) error {
	fmt.Println("Car deleted")
	return nil
}
