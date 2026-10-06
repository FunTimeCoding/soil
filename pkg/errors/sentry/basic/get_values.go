package basic

import (
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/url"
)

func (c *Client) GetValues(
	path string,
	query url.Values,
	out any,
) error {
	return c.requester.Notation(request.Get(path).WithParameters(query), out)
}
