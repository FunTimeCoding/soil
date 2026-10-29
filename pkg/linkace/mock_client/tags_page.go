package mock_client

import (
	"github.com/funtimecoding/soil/pkg/linkace/page"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

func (c *Client) TagsPage(_ int) (*page.Page[*tag.Tag], error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &page.Page[*tag.Tag]{LastPage: 1}, nil
}
