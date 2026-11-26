package ladybird

func (c *Client) forget(identifier int) {
	c.mutex.Lock()
	delete(c.pending, identifier)
	c.mutex.Unlock()
}
