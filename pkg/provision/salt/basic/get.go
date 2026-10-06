package basic

import "github.com/funtimecoding/soil/pkg/web/requester/request"

func (c *Client) Get(
	path string,
	out any,
) error {
	return c.requester.Notation(request.Get(path), out)
}
