package basic

import "github.com/funtimecoding/soil/pkg/web/requester/request"

func (c *Client) Bytes(l string) ([]byte, error) {
	return c.requester.Bytes(request.Absolute(l))
}
