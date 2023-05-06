package main

import (
	"main/http"
	"main/parser"
)

func main() {
	c := http.NewClient()
	crawler := parser.NewCrawler(c)

	crawler.Crawl("https://golangbot.com")
}
