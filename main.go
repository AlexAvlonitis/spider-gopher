package main

import (
	"main/http"
	"main/traverser"
)

func main() {
	c := http.NewClient()
	bfs := traverser.New(c)

	bfs.Traverse("https://alex.avlonitis.me")
}
