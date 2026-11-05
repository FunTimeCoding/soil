package mock_client

import "maps"

func (c *Client) AddSecret(
	path string,
	values map[string]string,
) {
	c.secrets[path] = maps.Clone(values)
}
