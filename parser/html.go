package parser

import (
	"bytes"
	"fmt"
	"net/url"

	"golang.org/x/net/html"
)

// Input a main path string URL and a stream of bytes from a parsed HTML document,
// and returns a list of strings containing all the hyperlink URLs of the same domain
func ExtractAllDomainLinks(m string, b []byte) []string {
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
				links = getHrefAttrs(m, t.Attr)
			}
		}
	}
}

// Get a list of href string from a html.Attribute node and return a list of strings URLs
// as long as in the same domain as the main path
func getHrefAttrs(mainPath string, attrs []html.Attribute) []string {
	var links []string

	for _, attr := range attrs {
		if attr.Key == "href" {
			if isSameDomain(mainPath, attr.Val) {
				links = append(links, attr.Val)
			}
		}
	}
	return links
}

// Compares two string URLs if they have the same host value
func isSameDomain(m, a string) bool {
	mainPath, err := url.Parse(m)
	if err != nil {
		return false
	}
	url, err := url.Parse(a)
	if err != nil {
		return false
	}

	fmt.Println(mainPath.Host + ":" + url.Host)
	return mainPath.Host == url.Host
}
