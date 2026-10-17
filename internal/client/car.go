package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rodolfodiazr/cars-cli/internal/car"
)

var (
	// ErrEmptyCarID indicates that an empty car ID was provided.
	ErrEmptyCarID = errors.New("empty car ID")
	// ErrCarNotFound indicates that the requested car could not be found.
	ErrCarNotFound           = errors.New("car not found")
	ErrCarMakeIsRequired     = errors.New("make is required")
	ErrCarModelIsRequired    = errors.New("model is required")
	ErrCarColorIsRequired    = errors.New("color is required")
	ErrCarCategoryIsRequired = errors.New("category is required")
	ErrCarInvalidYear        = errors.New("year is not valid")
	ErrUnknown               = errors.New("unknown error")
)

// Car defines the operations available for interacting with the Cars API.
type Car interface {
	ListCars(ctx context.Context, make, model string, year int) (car.CarsResponse, error)
	GetCar(ctx context.Context, carID string) (car.CarResponse, error)
	CreateCar(ctx context.Context, req car.CreateCar) (string, error)
	UpdateCar(ctx context.Context, carID string, req car.UpdateCar) error
	DeleteCar(ctx context.Context, carID string) error
}

type carClient struct {
	httpClient HttpClient
}

// NewCarClient creates a Car client that uses httpClient to send HTTP requests.
func NewCarClient(httpClient *http.Client) Car {
	return &carClient{
		httpClient: httpClient,
	}
}

// ListCars retrieves cars matching the specified filters.
func (c *carClient) ListCars(ctx context.Context, make, model string, year int) (car.CarsResponse, error) {
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
		return car.CarsResponse{}, err
	}
	request = request.WithContext(ctx)

	// Make the request
	var carsResponse CarsResponse
	if err := Do(c.httpClient, request, &carsResponse); err != nil {
		return nil, err
	}

	cars := car.CarsResponse{}
	for _, c := range carsResponse {
		cars = append(cars, car.CarResponse{
			ID:       c.CarID,
			Make:     c.Make,
			Model:    c.Model,
			Color:    c.Color,
			Category: c.Category,
			Year:     c.Year,
		})
	}
	return cars, nil
}

// GetCar retrieves a car by its ID.
func (c *carClient) GetCar(ctx context.Context, carID string) (car.CarResponse, error) {
	if carID == "" {
		return car.CarResponse{}, ErrEmptyCarID
	}
	path := fmt.Sprintf("/cars/%s", carID)

	request, err := BuildRequest(http.MethodGet, "http://localhost:8080", path, nil, nil)
	if err != nil {
		return car.CarResponse{}, err
	}
	request = request.WithContext(ctx)

	// Make the request
	var carResponse CarResponse
	if err := Do(c.httpClient, request, &carResponse); err != nil {
		var responseErr HttpResponseError
		if errors.As(err, &responseErr) {
			if responseErr.StatusCode == http.StatusNotFound {
				return car.CarResponse{}, ErrCarNotFound
			}
		}
		return car.CarResponse{}, err
	}

	return car.CarResponse{
		ID:       carResponse.CarID,
		Make:     carResponse.Make,
		Model:    carResponse.Model,
		Color:    carResponse.Color,
		Category: carResponse.Category,
		Year:     carResponse.Year,
	}, nil
}

// CreateCar inserts a new car.
func (c *carClient) CreateCar(ctx context.Context, req car.CreateCar) (string, error) {
	path := "/cars"

	payload := CreateCarRequest{
		Make:     req.Make,
		Model:    req.Model,
		Color:    req.Color,
		Category: req.Category,
		Year:     req.Year,
	}

	request, err := BuildRequest(http.MethodPost, "http://localhost:8080", path, nil, payload)
	if err != nil {
		return "", err
	}
	request = request.WithContext(ctx)

	// Make the request
	var response CarResponse
	if err := Do(c.httpClient, request, &response); err != nil {
		var responseErr HttpResponseError
		if errors.As(err, &responseErr) {
			if responseErr.StatusCode == http.StatusBadRequest {
				var validationErr CarValidationError
				if err := json.Unmarshal([]byte(responseErr.ResponseBody), &validationErr); err != nil {
					return "", err
				}
				return "", c.translateValidationError(validationErr)
			}
		}
		return "", err
	}
	return response.CarID, nil
}

// UpdateCar updates a car by its id.
func (c *carClient) UpdateCar(ctx context.Context, carID string, req car.UpdateCar) error {
	if carID == "" {
		return ErrEmptyCarID
	}
	path := fmt.Sprintf("/cars/%s", carID)

	payload := UpdateCarRequest{
		Make:     req.Make,
		Model:    req.Model,
		Color:    req.Color,
		Category: req.Category,
		Year:     req.Year,
	}

	request, err := BuildRequest(http.MethodPut, "http://localhost:8080", path, nil, payload)
	if err != nil {
		return err
	}
	request = request.WithContext(ctx)

	// Make the request
	if err := Do(c.httpClient, request, nil); err != nil {
		var responseErr HttpResponseError
		if errors.As(err, &responseErr) {
			if responseErr.StatusCode == http.StatusBadRequest {
				var validationErr CarValidationError
				if err := json.Unmarshal([]byte(responseErr.ResponseBody), &validationErr); err != nil {
					return err
				}
				return c.translateValidationError(validationErr)
			}
		}
		return err
	}
	return nil
}

func (c *carClient) DeleteCar(ctx context.Context, carID string) error {
	if carID == "" {
		return ErrEmptyCarID
	}
	path := fmt.Sprintf("/cars/%s", carID)

	request, err := BuildRequest(http.MethodDelete, "http://localhost:8080", path, nil, nil)
	if err != nil {
		return err
	}
	request = request.WithContext(ctx)

	// Make the request
	if err := Do(c.httpClient, request, nil); err != nil {
		var responseErr HttpResponseError
		if errors.As(err, &responseErr) {
			var responseErr HttpResponseError
			if errors.As(err, &responseErr) {
				if responseErr.StatusCode == http.StatusNotFound {
					return ErrCarNotFound
				}
			}
			return err
		}
		return err
	}
	return nil
}

func (c *carClient) translateValidationError(validationErr CarValidationError) error {
	if validationErr.Details == "make is required" {
		return ErrCarMakeIsRequired
	}
	if validationErr.Details == "model is required" {
		return ErrCarModelIsRequired
	}
	if validationErr.Details == "color is required" {
		return ErrCarColorIsRequired
	}
	if validationErr.Details == "category is required" {
		return ErrCarCategoryIsRequired
	}
	if validationErr.Details == "year is not valid" {
		return ErrCarInvalidYear
	}
	return ErrUnknown
}
