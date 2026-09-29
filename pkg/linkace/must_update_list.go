package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/list"
)

func (c *Client) MustUpdateList(
	identifier int,
	patch map[string]any,
) *list.List {
	result, e := c.UpdateList(identifier, patch)
	errors.PanicOnError(e)

	return result
}
