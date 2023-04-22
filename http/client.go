package http

type Client interface {
	GetResponseBody(string) ([]byte, error)
}
