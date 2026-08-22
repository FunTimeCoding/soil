package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) SaveView(
	identifier int,
	path string,
) error {
	for i, v := range c.views {
		if v.Identifier == identifier {
			c.views[i].FilePath = path
			c.views[i].Dirty = false

			return nil
		}
	}

	return not_found.New("view", identifier)
}
