package mock_client

import "github.com/funtimecoding/soil/pkg/sublime/view"

func (c *Client) Views() ([]*view.View, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.views, nil
}
