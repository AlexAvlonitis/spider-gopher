package main

func main() {
	c := NewHttpClient()
	bst := NewBfsTraverser(c)

	bst.traverse("https://alex.avlonitis.me")
}
