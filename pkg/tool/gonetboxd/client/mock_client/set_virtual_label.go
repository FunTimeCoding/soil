package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) SetVirtualLabel(_ context.Context, _ string, _ string, _ string) (client.Label, error) {
	return client.Label{}, c.Failure
}
