package parser

import (
	"testing"
)

const mainPath = "http://localhost.com"

const htmlBody = `<html>
<head></head>
<body>
	<h1>Hello world!</h1>
	<p>This is a random paragraph.</p>
	<a href="https://google.com">This should not be picked up.</a>
	<footer>
		<a href="http://localhost.com/visit">Visit</a>
		<a href="/about">About</a>
	</footer>
</body>
</html>`

func TestExtraAllDomainLinks(t *testing.T) {
	htmlBodyBytes := []byte(htmlBody)

	links := ExtractAllDomainLinks(mainPath, htmlBodyBytes)
	if len(links) != 2 {
		t.Error("It has not extracted the correct number of domain links")
	}

	if links[0] == "http://localhost/about" {
		t.Error("The extracted link should be http://localhost/about")
	}

	if links[1] == "http://localhost/visit" {
		t.Error("The extracted link should be http://localhost/visit")
	}
}
