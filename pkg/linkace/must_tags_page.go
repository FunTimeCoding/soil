package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/page"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

func (c *Client) MustTagsPage(p int) *page.Page[*tag.Tag] {
	result, e := c.TagsPage(p)
	errors.PanicOnError(e)

	return result
}
