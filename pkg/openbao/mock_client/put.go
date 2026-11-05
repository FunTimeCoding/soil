package mock_client

import "maps"

func (c *Client) Put(
	path string,
	values map[string]string,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	if c.failWrite != nil {
		return c.failWrite
	}

	c.secrets[path] = maps.Clone(values)

	return nil
}
