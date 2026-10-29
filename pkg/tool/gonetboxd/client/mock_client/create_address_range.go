package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateAddressRange(_ context.Context, _ string, _ string, _ string, _ string) (client.AddressRange, error) {
	return client.AddressRange{}, c.Failure
}
