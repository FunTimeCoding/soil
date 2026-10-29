package mock_client

import (
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/linkace/page"
)

func (c *Client) SearchPage(
	_ string,
	_ int,
) (*page.Page[*link.Link], error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &page.Page[*link.Link]{LastPage: 1}, nil
}
