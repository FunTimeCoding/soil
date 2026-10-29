package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListAddressRanges(_ context.Context) ([]*client.AddressRange, error) {
	return nil, c.Failure
}
