package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chromium/snapshot"
)

func (c *Client) Snapshot(_ context.Context) ([]*snapshot.Node, error) {
	return nil, nil
}
