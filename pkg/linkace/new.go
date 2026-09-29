package linkace

import "github.com/funtimecoding/soil/pkg/linkace/basic"

func New(
	host string,
	token string,
) *Client {
	return &Client{basic: basic.New(host, token), host: host}
}
