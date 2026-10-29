package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateManufacturer(_ context.Context, _ string) (client.Manufacturer, error) {
	return client.Manufacturer{}, c.Failure
}
