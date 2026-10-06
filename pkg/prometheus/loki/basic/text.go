package basic

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Text(l string) (string, error) {
	if c.verbose {
		console.Format("GET %s\n", l)
	}

	b, e := c.requester.Bytes(request.Absolute(l))

	return string(b), e
}
