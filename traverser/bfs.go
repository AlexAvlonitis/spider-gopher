package traverser

import (
	"fmt"
	"log"
	"main/dts"
	"main/http"
	"main/parser"
)

type Bfs struct {
	client       http.Client
	visitedLinks map[string]bool
	queue        dts.Queue
}

func (b *Bfs) Traverse(path string) {
	b.queue.Enqueue(path)

	for b.queue.Size() > 0 {
		link := b.queue.NextValue()
		b.queue.Dequeue()
		fmt.Println(link)

		respBody, err := b.client.GetResponseBody(link)
		if err != nil {
			log.Fatal(err)
		}

		links := parser.ExtractAllDomainLinks(path, respBody)
		for _, l := range links {
			if !b.visitedLinks[l] {
				b.visitedLinks[l] = true
				b.queue.Enqueue(l)
			}
		}
	}
}

func New(c http.Client) *Bfs {
	return &Bfs{
		client:       c,
		visitedLinks: make(map[string]bool),
		queue:        dts.NewQueue(),
	}
}
