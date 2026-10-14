package mock_client

func (c *Client) WithPodError(e error) *Client {
	c.podError = e

	return c
}
