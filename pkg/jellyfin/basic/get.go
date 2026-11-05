package basic

import (
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/url"
)

func (c *Client) Get(
	path string,
	v url.Values,
	out any,
) error {
	return c.requester.Notation(
		c.classified(request.Get(path).WithParameters(v)),
		out,
	)
}
