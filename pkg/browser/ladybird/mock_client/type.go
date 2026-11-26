package mock_client

func (c *Client) Type(
	context string,
	text string,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, e := c.find(context); e != nil {
		return e
	}

	c.typed = append(c.typed, text)

	return nil
}
