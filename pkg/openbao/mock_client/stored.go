package mock_client

import "maps"

func (c *Client) Stored(path string) map[string]string {
	return maps.Clone(c.secrets[path])
}
