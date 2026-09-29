package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/link"
)

func (c *Client) MustSearch(query string) []*link.Link {
	result, e := c.Search(query)
	errors.PanicOnError(e)

	return result
}
