package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/link"
)

func (c *Client) MustLinks() []*link.Link {
	result, e := c.Links()
	errors.PanicOnError(e)

	return result
}
