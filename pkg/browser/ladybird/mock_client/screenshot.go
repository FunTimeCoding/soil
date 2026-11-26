package mock_client

func (c *Client) Screenshot(context string) ([]byte, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, e := c.find(context); e != nil {
		return nil, e
	}

	return c.image, nil
}
