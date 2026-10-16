package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// HttpResponseError represents an HTTP response with a status code indicating
// a client or server error. It includes the request URL and response body
// to provide additional context about the failure.
type HttpResponseError struct {
	StatusCode   int
	RequestURL   *url.URL
	ResponseBody string
}

// Error implements the error interface.
func (e HttpResponseError) Error() string {
	return fmt.Sprintf("Status %d: Request URL: %v Error Body: %s", e.StatusCode, e.RequestURL, e.ResponseBody)
}

// HttpClient defines a HTTP client that supports mocking.
//
// Use this when you need direct control over the raw *http.Response,
// including non-JSON responses or custom handling logic.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// BuildRequest builds the http request based on the input parameters
func BuildRequest(method, baseURL, path string, query url.Values, body any) (*http.Request, error) {
	endpoint := strings.TrimRight(baseURL, "/") + path

	if len(query) > 0 {
		endpoint += "?" + query.Encode()
	}

	var reader io.Reader
	if body != nil {
		r, err := serializePayload(body)
		if err != nil {
			return nil, fmt.Errorf("failed to serialize request body: %w", err)
		}
		reader = r
	}

	req, err := http.NewRequest(method, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do executes an HTTP request using the provided client and
// handles standard JSON request/response processing.
func Do(httpClient HttpClient, request *http.Request, v any) error {
	resp, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read raw body once
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Handle errors
	if resp.StatusCode >= http.StatusBadRequest {
		return HttpResponseError{
			StatusCode:   resp.StatusCode,
			RequestURL:   request.URL,
			ResponseBody: string(rawBody),
		}
	}

	// Success decode
	if v != nil && len(rawBody) > 0 {
		if err := json.Unmarshal(rawBody, v); err != nil {
			return err
		}
		fmt.Println("v: ", v)
	}

	return nil
}

// serializePayload converts a request payload into an io.Reader suitable for an HTTP request body.
//
// It supporst the following payload types:
// - url.Values or *url.Values: encoded as application/x-www-form-urlencoded
// - []byte: used directly as the request body
// - io.Reader: used as-is
//
// For any other type, the payload is JSON-marshaled.
// An error is returned if the payload cannot be serialized.
func serializePayload(body any) (io.Reader, error) {
	switch b := body.(type) {
	case url.Values:
		return strings.NewReader(b.Encode()), nil

	case *url.Values:
		return strings.NewReader(b.Encode()), nil

	case []byte:
		return bytes.NewReader(b), nil

	case io.Reader:
		return b, nil

	default:
		buf, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		return bytes.NewBuffer(buf), nil
	}
}
