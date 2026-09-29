package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

func (c *Client) MustUpdateTag(
	identifier int,
	patch map[string]any,
) *tag.Tag {
	result, e := c.UpdateTag(identifier, patch)
	errors.PanicOnError(e)

	return result
}
