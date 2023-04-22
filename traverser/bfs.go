package traverser

import (
	"fmt"
	"log"
	"main/dts"
	"main/http"
	"main/parser"
)

type Bfs struct {
	client http.Client
	set    dts.Set
	queue  dts.Queue
}

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

func New(c http.Client) Traverser {
	return &Bfs{
		client: c,
		set:    dts.NewSet(),
		queue:  dts.NewQueue(),
	}
}
