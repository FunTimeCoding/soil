package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/linkace/page"
)

func (c *Client) MustLinksPage(p int) *page.Page[*link.Link] {
	result, e := c.LinksPage(p)
	errors.PanicOnError(e)

	return result
}
