package basic

import "github.com/funtimecoding/soil/pkg/web/requester/request"

func (c *Client) Get(
	l string,
	out any,
) error {
	return c.requester.Notation(request.Absolute(l), out)
}
