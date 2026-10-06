package mock_client

func (c *Client) FailPipelines(e error) {
	c.pipelineFailure = e
}
