package http

import (
	"io"
	"net/http"
)

type DefaultClient struct {
	client *http.Client
}

type ResponseUrl struct {
	Path     string
	HtmlBody []byte
}

// HTTP Get request, returns the parsed html response as a ResponseUrl object
func (d *DefaultClient) GetResponse(path string) (ResponseUrl, error) {
	resp, err := d.client.Get(path)
	if err != nil {
		return ResponseUrl{}, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseUrl{}, err
	}

	l := ResponseUrl{Path: path, HtmlBody: respBody}
	return l, nil
}

func NewClient() Client {
	return &DefaultClient{&http.Client{}}
}
