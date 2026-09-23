package chromium

func (c *Client) TargetCount() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return len(c.targets)
}
