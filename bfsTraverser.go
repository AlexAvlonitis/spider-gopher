package main

import (
	"fmt"
	"log"
)

type BfsTraverser struct {
	client       httpClient
	visitedLinks map[string]bool
	queue        Queue
}

func (b *BfsTraverser) traverse(path string) {
	b.queue.Enqueue(path)

	for b.queue.Size() > 0 {
		link := b.queue.ExitValue()
		b.queue.Dequeue()
		fmt.Println(link)

		respBody, err := b.client.GetResponseBody(link)
		if err != nil {
			log.Fatal(err)
		}

		links := ExtractAllLinks(respBody)
		for _, l := range links {
			if !b.visitedLinks[l] {
				b.visitedLinks[l] = true
				b.queue.Enqueue(l)
			}
		}
	}
}

func NewBfsTraverser(c httpClient) *BfsTraverser {
	return &BfsTraverser{
		client:       c,
		visitedLinks: make(map[string]bool),
		queue:        NewQueue(),
	}
}
