package mock_client

func (c *Client) Fail(
	method any,
	e error,
) {
	c.failures.Set(method, e)
}
