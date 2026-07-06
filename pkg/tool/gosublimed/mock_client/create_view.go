package mock_client

import "github.com/funtimecoding/soil/pkg/sublime/view"

func (c *Client) CreateView(
	title string,
	content string,
	syntax string,
) (*view.View, error) {
	v := view.Stub()
	v.Identifier = c.nextID
	v.Title = title
	v.Text = content
	v.Syntax = syntax
	c.nextID++
	c.views = append(c.views, v)

	return v, nil
}
