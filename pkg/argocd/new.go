package argocd

import "github.com/funtimecoding/go-library/pkg/web"

func New(
	host string,
	port int,
	secure bool,
	token string,
) *Client {
	return &Client{
		base:  web.Link(host, port, secure),
		token: token,
	}
}
