package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"slices"
)

func (c *Client) ListFindings(context.Context) ([]*client.FindingGroup, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	return slices.Clone(c.groups), nil
}
