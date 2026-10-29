package mock_client

import "context"

func (c *Client) TransitionIssue(
	_ context.Context,
	_ string,
	_ string,
) error {
	return c.Failure
}
