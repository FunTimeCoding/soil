package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/list"
)

func (c *Client) MustCreateList(
	name string,
	description string,
) *list.List {
	result, e := c.CreateList(name, description)
	errors.PanicOnError(e)

	return result
}
