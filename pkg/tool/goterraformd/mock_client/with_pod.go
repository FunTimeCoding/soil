package mock_client

func (c *Client) WithPod(name string) *Client {
	c.pods[name] = true

	return c
}
