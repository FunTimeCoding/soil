package client

import "github.com/funtimecoding/soil/pkg/web/stall_client"

func New(
	base string,
	untrusted bool,
	token string,
) *Client {
	return newOnPool(stall_client.New().Traced(),base, untrusted, token)
}
