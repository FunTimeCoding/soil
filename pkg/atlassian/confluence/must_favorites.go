package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustFavorites() []*page.Page {
	result, e := c.Favorites()
	errors.PanicOnError(e)

	return result
}
