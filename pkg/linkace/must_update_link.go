package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/link"
)

func (c *Client) MustUpdateLink(
	identifier int,
	patch map[string]any,
) *link.Link {
	result, e := c.UpdateLink(identifier, patch)
	errors.PanicOnError(e)

	return result
}
