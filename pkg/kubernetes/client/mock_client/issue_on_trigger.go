package mock_client

func (c *Client) IssueOnTrigger() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.issueOnTrigger = true
}
