package mock_client

func (c *Client) WriteFile(
	path string,
	content []byte,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return e
	}

	c.files[path] = content

	return nil
}
