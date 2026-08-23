package telemetry

func (c *Client) Flush() {
	c.group.Wait()
}
