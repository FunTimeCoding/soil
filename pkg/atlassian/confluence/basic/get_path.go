package basic

import "github.com/funtimecoding/soil/pkg/web/requester/request"

func (c *Client) GetPath(
	path string,
	out any,
) error {
	return c.old.Notation(request.Get(path), out)
}
