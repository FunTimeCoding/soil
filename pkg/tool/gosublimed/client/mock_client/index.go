package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) index(identifier int) (int, error) {
	for i, v := range c.ViewList {
		if v.ViewId == identifier {
			return i, nil
		}
	}

	return 0, not_found.New("view", identifier)
}
