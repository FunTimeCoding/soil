package mock_client

import "github.com/funtimecoding/soil/pkg/sublime/view"

func (c *Client) Views() ([]*view.View, error) {
	return c.views, nil
}
