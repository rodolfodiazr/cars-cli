package client

import (
	"context"
	"net/http"

	"github.com/rodolfodiazr/cars-cli/internal/car"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
}

func New(httpClient *http.Client, baseURL string) *Client {
	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

func (c *Client) ListCars(ctx context.Context) ([]car.Car, error) {
	// GET /cars
	return []car.Car{}, nil
}

func (c *Client) GetCar(ctx context.Context, id string) (car.Car, error) {
	// GET /cars/{id}
	return car.Car{}, nil
}

func (c *Client) CreateCar(ctx context.Context, car car.Car) error {
	// POST /cars
	return nil
}

func (c *Client) UpdateCar(ctx context.Context, car car.Car) error {
	// PUT/PATCH /cars/{id}
	return nil
}

func (c *Client) DeleteCar(ctx context.Context, id string) error {
	// DELETE /cars/{id}
	return nil
}
