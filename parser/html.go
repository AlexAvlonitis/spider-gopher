package parser

import (
	"bytes"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// Input a main path string URL and a stream of bytes from a parsed HTML document,
// and returns a list of strings containing all the hyperlink URLs of the same domain
func ExtractAllDomainLinks(mainPath string, b []byte) []string {
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
				links = append(links, getHrefAttrs(mainPath, t.Attr)...)
			}
		}
	}
}

// Returns the href attribute url. If the href attribute is not an absolute path
// it appends the full path.
func getHrefAttrs(mainPath string, attrs []html.Attribute) []string {
	var links []string

	for _, attr := range attrs {
		if attr.Key == "href" {
			if isSameWebsite(mainPath, attr.Val) {
				links = append(links, appendHost(mainPath, attr.Val))
			}
		}
	}
	return links
}

// Compares two string URLs if they are on the same website.
// Returns true if they have the same domain or a relative path
func isSameWebsite(m, u string) bool {
	mainPath, err := url.Parse(m)
	if err != nil {
		return false
	}
	url, err := url.Parse(u)
	if err != nil {
		return false
	}
	if url.IsAbs() {
		return (mainPath.Host == url.Host)
	}

	return strings.HasPrefix(url.Path, "/")
}

// Appends the scheme and host name of the main url to the link,
// if the url is already not absolute
func appendHost(mainPath, link string) string {
	m, err := url.Parse(mainPath)
	if err != nil {
		panic(err)
	}
	url, err := url.Parse(link)
	if err != nil {
		panic(err)
	}
	if url.IsAbs() {
		return link
	}

	return m.Scheme + "://" + m.Host + link
}
