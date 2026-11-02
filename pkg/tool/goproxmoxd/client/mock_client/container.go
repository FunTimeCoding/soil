package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) container(identifier int64) (int, error) {
	if c.Failure != nil {
		return 0, c.Failure
	}

	for i, v := range c.ContainerList {
		if v.Identifier == identifier {
			return i, nil
		}
	}

	return 0, not_found.Format("container %d", identifier)
}
