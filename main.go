package main

import (
	"fmt"
	"main/http"
	"main/parser"
)

func main() {
	c := http.NewClient()
	crawler := parser.NewCrawler(c)

	fmt.Println("Crawling, please wait...")
	crawler.Crawl("https://golangbot.com")

	for _, r := range crawler.Results {
		fmt.Println(r)
	}
}
