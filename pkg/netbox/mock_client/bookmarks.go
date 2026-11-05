package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/bookmark"

func (c *Client) Bookmarks() ([]*bookmark.Bookmark, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
