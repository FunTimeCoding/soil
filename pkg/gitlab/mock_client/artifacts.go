package mock_client

import (
	"bytes"
	"context"
)

func (c *Client) Artifacts(
	_ context.Context,
	_ int64,
	_ int64,
) (*bytes.Reader, error) {
	return nil, c.failure
}
