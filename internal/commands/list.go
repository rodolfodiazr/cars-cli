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
