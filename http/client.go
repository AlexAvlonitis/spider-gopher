package http

type Client interface {
	GetResponse(string) (Link, error)
}
