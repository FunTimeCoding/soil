package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateAddress(_ context.Context, _ string, _ string, _ string, _ string) (client.Address, error) {
	return client.Address{}, c.Failure
}
