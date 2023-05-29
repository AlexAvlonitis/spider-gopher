package http

type Client interface {
	GetResponse(string) (ResponseUrl, error)
}
