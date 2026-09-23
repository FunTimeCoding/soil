package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/mock_page"
)

func (c *Client) Page(_ string) face.Page {
	return mock_page.New()
}
