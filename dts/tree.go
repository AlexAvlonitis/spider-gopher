package dts

import (
	"fmt"
	"log"
	"main/http"
	"main/parser"
)

type Tree interface {
	Traverse(string)
}

// Simple BreadthFirst parser implementation
type Bfs struct {
	client http.Client
	set    Set
	queue  Queue
}

// Traverses all the links of a website in a breadth-first manner.
// Logs the already visited links to avoid circular requests.// Parse
func (b *Bfs) Traverse(path string) {
	b.queue.Enqueue(path)
	b.set.Add(path)

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
			if !b.set.Exists(l) {
				b.set.Add(l)
				b.queue.Enqueue(l)
			}
		}
	}
}

func NewBfs(c http.Client) Tree {
	return &Bfs{
		client: c,
		set:    NewSet(),
		queue:  NewQueue(),
	}
}
