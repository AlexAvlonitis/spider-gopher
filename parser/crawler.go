package parser

import (
	"fmt"
	"log"
	"main/dts"
	"main/http"
	"sync"
)

type Crawler struct {
	httpClient   http.Client
	links        dts.Queue
	visitedLinks dts.Set
}

// Traverses all the links of a website in a breadth-first iterative manner.
// Uses wait groups for concurrency, batches based on the number of links per page
func (c *Crawler) Crawl(path string) {
	c.links.Enqueue(path)
	c.visitedLinks.Add(path)

	for c.links.IsNotEmpty() {
		lks := c.links.DequeueAll()
		var wg sync.WaitGroup
		var mu sync.Mutex

		for _, l := range lks {
			wg.Add(1)
			go c.fetchLinks(l, &wg, &mu)
		}
		wg.Wait()
	}
}

// GET request for the given link, mutex locking the concurrent read/writes.
func (c *Crawler) fetchLinks(link string, wg *sync.WaitGroup, m *sync.Mutex) {
	l, err := c.httpClient.GetResponse(link)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(l.Path) // Print the path to stdout, TODO: refactor to a printer function

	links := ExtractAllDomainLinks(l.Path, l.HtmlBody)
	for _, l := range links {
		m.Lock()
		if !c.visitedLinks.Exists(l) {
			c.visitedLinks.Add(l)
			c.links.Enqueue(l)
		}
		m.Unlock()
	}
	wg.Done()
}

func NewCrawler(c http.Client) *Crawler {
	return &Crawler{
		httpClient:   c,
		visitedLinks: dts.NewSet(),
		links:        dts.NewQueue(),
	}
}
