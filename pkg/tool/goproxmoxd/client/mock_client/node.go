package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) node(name string) error {
	if c.Failure != nil {
		return c.Failure
	}

	for _, n := range c.NodeList {
		if n.Name == name {
			return nil
		}
	}

	return not_found.Format("node %s", name)
}
