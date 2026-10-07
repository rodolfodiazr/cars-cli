package client

import (
	"context"
	"encoding/json"
	"fmt"
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
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.baseURL+"/cars",
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var cars []car.Car
	if err := json.NewDecoder(resp.Body).Decode(&cars); err != nil {
		return nil, err
	}

	return cars, nil
}

func (c *Client) GetCar(ctx context.Context, id string) (car.Car, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.baseURL+"/cars/"+id,
		nil,
	)
	if err != nil {
		return car.Car{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return car.Car{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return car.Car{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var rs car.Car
	if err := json.NewDecoder(resp.Body).Decode(&rs); err != nil {
		return car.Car{}, err
	}

	return rs, nil
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
