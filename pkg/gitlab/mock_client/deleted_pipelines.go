package mock_client

func (c *Client) DeletedPipelines() []int64 {
	return c.deletedPipelines
}
