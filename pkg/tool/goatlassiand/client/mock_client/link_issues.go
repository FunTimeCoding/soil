package mock_client

import "context"

func (c *Client) LinkIssues(
	_ context.Context,
	_ string,
	_ string,
	_ string,
) error {
	return c.Failure
}
