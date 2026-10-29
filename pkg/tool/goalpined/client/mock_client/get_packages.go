package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/generated/client"
)

func (c *Client) GetPackages(
	_ context.Context,
	_ *client.GetPackagesParams,
) ([]client.Listing, error) {
	return nil, c.Failure
}
