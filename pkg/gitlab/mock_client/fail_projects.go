package mock_client

func (c *Client) FailProjects(e error) {
	c.projectFailure = e
}
