package client

import "github.com/funtimecoding/soil/pkg/web/stall_client"

func NewLong(
	base string,
	untrusted bool,
	token string,
) *Client {
	return newOnPool(stall_client.New().Long(), base, untrusted, token)
}
