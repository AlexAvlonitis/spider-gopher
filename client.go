package main

import (
	"io"
	"net/http"
)

type httpClient interface {
	GetResponseBody(string) ([]byte, error)
}

type DefaultHttpClient struct {
	client *http.Client
}

func (d DefaultHttpClient) GetResponseBody(path string) ([]byte, error) {
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

func NewHttpClient() *DefaultHttpClient {
	return &DefaultHttpClient{&http.Client{}}
}
