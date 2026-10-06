package basic

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Get(
	l string,
	out any,
) error {
	if c.verbose {
		console.Format("GET %s\n", l)
	}

	return c.requester.Notation(request.Absolute(l), out)
}
