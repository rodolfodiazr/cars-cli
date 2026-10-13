package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rodolfodiazr/cars-cli/internal/car"
)

var (
	ErrEmptyCarID  = errors.New("empty car ID")
	ErrCarNotFound = errors.New("car not found")
)

type Car interface {
	ListCars(ctx context.Context, make, model string, year int) ([]car.Car, error)
	GetCar(ctx context.Context, carID string) (car.Car, error)
	CreateCar(ctx context.Context, car car.Car) error
	UpdateCar(ctx context.Context, car car.Car) error
	DeleteCar(ctx context.Context, carID string) error
}

type carClient struct {
	httpClient HttpClient
}

func NewCarClient(httpClient *http.Client) Car {
	return &carClient{
		httpClient: httpClient,
	}
}

func (c *carClient) ListCars(ctx context.Context, make, model string, year int) ([]car.Car, error) {
	path := "/cars"

	// Set query params
	params := url.Values{}
	if make != "" {
		params.Add("make", make)
	}
	if model != "" {
		params.Add("model", model)
	}
	if year > 0 {
		params.Add("year", fmt.Sprintf("%d", year))
	}

	request, err := BuildRequest(http.MethodGet, "http://localhost:8080", path, params, nil)
	if err != nil {
		return []car.Car{}, err
	}
	request = request.WithContext(ctx)

	// Make the request
	var carsResponse car.CarsResponse
	if err := Do(c.httpClient, request, &carsResponse); err != nil {
		return nil, err
	}

	cars := []car.Car{}
	for _, c := range carsResponse {
		cars = append(cars, car.Car{
			ID:    c.ID,
			Make:  c.Make,
			Model: c.Model,
			Year:  c.Year,
		})
	}
	return cars, nil
}

func (c *carClient) GetCar(ctx context.Context, carID string) (car.Car, error) {
	if carID == "" {
		return car.Car{}, ErrEmptyCarID
	}

	path := fmt.Sprintf("/cars/%s", carID)

	request, err := BuildRequest(http.MethodGet, "http://localhost:8080", path, nil, nil)
	if err != nil {
		return car.Car{}, err
	}
	request = request.WithContext(ctx)

	// Make the request
	var carResponse car.CarResponse
	if err := Do(c.httpClient, request, &carResponse); err != nil {
		var responseErr car.CarResponseError
		if errors.As(err, &responseErr) {
			if responseErr.StatusCode == http.StatusNotFound {
				return car.Car{}, ErrCarNotFound
			}
		}
		return car.Car{}, err
	}

	return car.Car{
		ID:    carResponse.ID,
		Make:  carResponse.Make,
		Model: carResponse.Model,
		Year:  carResponse.Year,
	}, nil
}

func (c *carClient) CreateCar(ctx context.Context, car car.Car) error {
	// POST /cars
	return nil
}

func (c *carClient) UpdateCar(ctx context.Context, car car.Car) error {
	// url := fmt.Sprintf("%s/cars/%s", c.baseURL, car.ID)

	// body, err := json.Marshal(car)
	// if err != nil {
	// 	return err
	// }

	// req, err := http.NewRequestWithContext(
	// 	ctx,
	// 	http.MethodPut,
	// 	url,
	// 	bytes.NewReader(body),
	// )
	// if err != nil {
	// 	return err
	// }

	// req.Header.Set("Content-Type", "application/json")

	// resp, err := c.httpClient.Do(req)
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	// }

	return nil
}

func (c *carClient) DeleteCar(ctx context.Context, carID string) error {
	// DELETE /cars/{id}
	return nil
}
