package http

import (
	"io"
	"net/http"
)

type DefaultClient struct {
	client *http.Client
}

// HTTP Get request, returns the parsed html response in []bytes format.
func (d *DefaultClient) GetResponseBody(path string) ([]byte, error) {
	resp, err := d.client.Get(path)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func NewClient() Client {
	return &DefaultClient{&http.Client{}}
}
