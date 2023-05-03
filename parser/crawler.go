package parser

import (
	"fmt"
	"log"
	"main/dts"
	"main/http"
)

type Crawler struct {
	httpClient   http.Client
	links        dts.Queue
	visitedLinks dts.Set
}

// Traverses all the links of a website in a breadth-first iterative manner.
// Logs the already visited links to avoid circular requests.
func (b *Crawler) Crawl(path string) {
	b.links.Enqueue(path)
	b.visitedLinks.Add(path)

	for b.links.IsNotEmpty() {
		link := b.links.NextValue()
		b.links.Dequeue()
		fmt.Println(link)

		respBody, err := b.httpClient.GetResponseBody(link)
		if err != nil {
			log.Fatal(err)
		}

		links := ExtractAllDomainLinks(path, respBody)
		for _, l := range links {
			if !b.visitedLinks.Exists(l) {
				b.visitedLinks.Add(l)
				b.links.Enqueue(l)
			}
		}
	}
}

func NewCrawler(c http.Client) *Crawler {
	return &Crawler{
		httpClient:   c,
		visitedLinks: dts.NewSet(),
		links:        dts.NewQueue(),
	}
}
