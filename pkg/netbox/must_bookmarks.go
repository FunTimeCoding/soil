package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/bookmark"
)

func (c *Client) MustBookmarks() []*bookmark.Bookmark {
	result, e := c.Bookmarks()
	errors.PanicOnError(e)

	return result
}
