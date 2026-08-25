package mock_client

import "github.com/funtimecoding/soil/pkg/sublime/view"

func (c *Client) OpenFile(path string) (*view.View, error) {
	v := view.Stub()
	v.Identifier = c.nextIdentifier
	v.FilePath = path
	c.nextIdentifier++
	c.views = append(c.views, v)

	return v, nil
}
