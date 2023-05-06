package http

import (
	"io"
	"net/http"
)

type DefaultClient struct {
	client *http.Client
}

type Link struct {
	Path     string
	HtmlBody []byte
}

// HTTP Get request, returns the parsed html response as a Link object
func (d *DefaultClient) GetResponse(path string) (Link, error) {
	resp, err := d.client.Get(path)
	if err != nil {
		return Link{}, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Link{}, err
	}

	l := Link{Path: path, HtmlBody: respBody}
	return l, nil
}

func NewClient() Client {
	return &DefaultClient{&http.Client{}}
}
