package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/bookmark"

func (c *Client) Bookmarks() ([]*bookmark.Bookmark, error) {
	return nil, nil
}
