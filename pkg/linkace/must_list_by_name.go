package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/list"
)

func (c *Client) MustListByName(name string) *list.List {
	result, e := c.ListByName(name)
	errors.PanicOnError(e)

	return result
}
