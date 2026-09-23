package mock_client

func (c *Client) DeletePipeline(
	_ int64,
	identifier int64,
) error {
	c.deletedPipelines = append(c.deletedPipelines, identifier)

	return nil
}
