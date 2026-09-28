package mock_outpost_source

func (c *Client) Fail(e error) {
	c.fail = e
}
