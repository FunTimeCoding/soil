package mock_closer

func (c *Closer) Close() error {
	c.closed++

	return c.result
}
