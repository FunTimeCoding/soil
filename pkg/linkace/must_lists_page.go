package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/linkace/page"
)

func (c *Client) MustListsPage(p int) *page.Page[*list.List] {
	result, e := c.ListsPage(p)
	errors.PanicOnError(e)

	return result
}
