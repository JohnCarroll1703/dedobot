package telegram

import "net/http"

type Client struct {
	basePath string
	host     string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		client:   http.Client{},
		basePath: newBasePath(token),
	}
}

func newBasePath(token string) string {
	return "bot" + token
}
