package mock_client

import (
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/linkace/page"
)

func (c *Client) ListsPage(_ int) (*page.Page[*list.List], error) {
	return &page.Page[*list.List]{LastPage: 1}, nil
}
