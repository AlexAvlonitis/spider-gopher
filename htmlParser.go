package main

import (
	"bytes"

	"golang.org/x/net/html"
)

// Input a stream of bytes from a parsed HTML document,
// and returns a list of strings containing all the hyperlink URLs
func ParseLinks(b []byte) []string {
	tkn := html.NewTokenizer(bytes.NewReader(b))

	var links []string

	for {
		tt := tkn.Next()

		switch {
		case tt == html.ErrorToken:
			return links
		case tt == html.StartTagToken:
			t := tkn.Token()
			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}

	return links
}
