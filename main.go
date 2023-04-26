package main

import (
	"main/dts"
	"main/http"
)

func main() {
	c := http.NewClient()
	bfs := dts.NewBfs(c)

	bfs.Traverse("https://alex.avlonitis.me")
}
