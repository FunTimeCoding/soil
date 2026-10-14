package mock_closer

func (c *Closer) Closed() int {
	return c.closed
}
