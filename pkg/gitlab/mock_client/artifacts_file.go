package mock_client

import "context"

func (c *Client) ArtifactsFile(
	_ context.Context,
	_ int64,
	_ string,
	_ string,
) (string, error) {
	return "", c.failure
}
