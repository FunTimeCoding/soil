package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) CloseView(identifier int) error {
	for i, v := range c.views {
		if v.Identifier == identifier {
			c.views = append(c.views[:i], c.views[i+1:]...)

			return nil
		}
	}

	return not_found.New("view", identifier)
}
