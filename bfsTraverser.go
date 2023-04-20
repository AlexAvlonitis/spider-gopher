package main

import (
	"fmt"
	"log"
)

type BfsTraverser struct {
	client       httpClient
	visitedLinks map[string]bool
	queue        []string
}

func (b *BfsTraverser) traverse(path string) {
	b.queue = append(b.queue, path)

	for len(b.queue) > 0 {
		link := b.queue[0]
		b.queue = b.queue[1:]
		fmt.Println(link)

		respBody, err := b.client.GetResponseBody(link)
		if err != nil {
			log.Fatal(err)
		}

		links := ExtractAllLinks(respBody)
		for _, l := range links {
			if !b.visitedLinks[l] {
				b.visitedLinks[l] = true
				b.queue = append(b.queue, l)
			}
		}
	}
}

func NewBfsTraverser(c httpClient) *BfsTraverser {
	return &BfsTraverser{
		client:       c,
		visitedLinks: make(map[string]bool),
		queue:        make([]string, 0),
	}
}
