package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gochromed/face"

func (c *Client) SetPage(
	identifier string,
	p face.Page,
) {
	c.pages[identifier] = p
}
