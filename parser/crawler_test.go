package parser

import (
	"main/http"
	"testing"
)

func TestCrawl(t *testing.T) {
	httpClient := http.NewMockedClient()
	c := NewCrawler(httpClient)
	c.Crawl("http://localhost.com")
	r := c.Results

	if len(r) != 2 {
		t.Error("It has not extracted the correct number of domain links")
	}

	if r[0] == "http://localhost.com/about" {
		t.Error("The extracted link should be http://localhost.com/about")
	}

	if r[1] == "http://localhost.com" {
		t.Error("The extracted link should be http://localhost.com")
	}
}
