package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/link"
)

func (c *Client) MustLinkByIdentifier(identifier int) *link.Link {
	result, e := c.LinkByIdentifier(identifier)
	errors.PanicOnError(e)

	return result
}
