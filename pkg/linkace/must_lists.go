package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/list"
)

func (c *Client) MustLists() []*list.List {
	result, e := c.Lists()
	errors.PanicOnError(e)

	return result
}
