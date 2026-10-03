package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/mock_page"
)

func (c *Client) Page(identifier string) face.Page {
	if p, found := c.pages[identifier]; found {
		return p
	}

	return mock_page.New()
}
