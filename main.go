package main

import (
	"fmt"
)

func main() {
	c := NewHttpClient()

	respBody, err := c.GetResponseBody("https://alex.avlonitis.me")
	if err != nil {
		panic(err)
	}

	links := ParseLinks(respBody)
	for _, val := range links {
		fmt.Println(val)
	}
}
