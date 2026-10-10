package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient HttpClient
	baseURL    string
}

func New(httpClient HttpClient, baseURL string) *Client {
	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

func (c *Client) Do(ctx context.Context, method, url string, body any) (*http.Response, error) {
	var requestBody *bytes.Buffer

	if body != nil {
		requestBody = new(bytes.Buffer)

		if err := json.NewEncoder(requestBody).Encode(body); err != nil {
			return nil, fmt.Errorf("encode request body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		url,
		requestBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}

	return resp, nil
}
