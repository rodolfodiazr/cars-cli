package client

import (
	"context"
	"net/http"

	"github.com/rodolfodiazr/cars-cli/internal/car"
)

type Car interface {
	ListCars(ctx context.Context) ([]car.Car, error)
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

func (c *carClient) ListCars(ctx context.Context) ([]car.Car, error) {
	path := "/cars"

	request, err := BuildRequest(http.MethodGet, "http://localhost:8080", path, nil, nil)
	if err != nil {
		return []car.Car{}, err
	}
	request = request.WithContext(ctx)

	// Make the request
	var carsResponse car.CarsResponse
	if err := Do(c.httpClient, request, &carsResponse); err != nil {
		// Handle known errors
		// var responseErr car.CarResponseError
		// if errors.As(err, &responseErr) {

		// }
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
	// req, err := http.NewRequestWithContext(
	// 	ctx,
	// 	http.MethodGet,
	// 	c.baseURL+"/cars/"+id,
	// 	nil,
	// )
	// if err != nil {
	// 	return car.Car{}, err
	// }

	// resp, err := c.httpClient.Do(req)
	// if err != nil {
	// 	return car.Car{}, err
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	return car.Car{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	// }

	// var rs car.Car
	// if err := json.NewDecoder(resp.Body).Decode(&rs); err != nil {
	// 	return car.Car{}, err
	// }

	// return rs, nil
	return car.Car{}, nil
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
