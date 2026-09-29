package mock_client

import (
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/linkace/page"
)

func (c *Client) LinksPage(_ int) (*page.Page[*link.Link], error) {
	return &page.Page[*link.Link]{LastPage: 1}, nil
}
