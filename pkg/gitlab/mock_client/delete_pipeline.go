package mock_client

import "context"

func (c *Client) DeletePipeline(
	_ context.Context,
	_ int64,
	identifier int64,
) error {
	c.deletedPipelines = append(c.deletedPipelines, identifier)

	return nil
}
