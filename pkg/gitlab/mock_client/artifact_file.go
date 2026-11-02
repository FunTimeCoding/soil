package mock_client

import (
	"bytes"
	"context"
)

func (c *Client) ArtifactFile(
	_ context.Context,
	_ int64,
	_ int64,
	_ string,
) (*bytes.Reader, error) {
	return nil, c.failure
}
