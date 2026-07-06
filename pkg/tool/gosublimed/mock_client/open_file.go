package mock_client

import "github.com/funtimecoding/soil/pkg/sublime/view"

func (c *Client) OpenFile(path string) (*view.View, error) {
	v := view.Stub()
	v.Identifier = c.nextID
	v.FilePath = path
	c.nextID++
	c.views = append(c.views, v)

	return v, nil
}
