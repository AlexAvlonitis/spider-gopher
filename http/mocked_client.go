package http

import (
	"fmt"
)

type MockedClient struct{}

const mainPath = "http://localhost.com"
const aboutPath = "http://localhost.com/about"

const htmlBody = `<html>
<head></head>
<body>
	<h1>Hello world!</h1>
	<p>This is a random paragraph.</p>
	<a href="https://google.com">This should not be picked up.</a>
	<footer>
		<a href="/about">About</a>
	</footer>
</body>
</html>`

const aboutPage = `<html
<head>Abount</head>
<body>
	<h1>This is the About page</h1>
</body>
</html>`

// Mocks the http client with fixed response
func (d *MockedClient) GetResponse(path string) (ResponseUrl, error) {
	if path == "http://localhost.com" {
		return ResponseUrl{Path: mainPath, HtmlBody: []byte(htmlBody)}, nil
	}
	if path == "http://localhost.com/about" {
		return ResponseUrl{Path: aboutPath, HtmlBody: []byte(aboutPage)}, nil
	}

	return ResponseUrl{}, fmt.Errorf("Error")
}

func NewMockedClient() Client {
	return &MockedClient{}
}
