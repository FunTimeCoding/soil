package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/link"
)

func (c *Client) MustLinksByList(identifier int) []*link.Link {
	result, e := c.LinksByList(identifier)
	errors.PanicOnError(e)

	return result
}
