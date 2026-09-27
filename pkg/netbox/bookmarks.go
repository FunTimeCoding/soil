package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/bookmark"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func (c *Client) Bookmarks() ([]*bookmark.Bookmark, error) {
	result, _, e := c.client.ExtrasAPI.ExtrasBookmarksList(c.context).Limit(
		constant.PageLimit,
	).Execute()

	if e != nil {
		return nil, e
	}

	return bookmark.NewSlice(result.Results), nil
}
