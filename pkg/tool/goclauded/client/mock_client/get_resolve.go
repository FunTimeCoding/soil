package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetResolve(_ context.Context, _ *client.GetResolveParams) (*client.ResolveResponse, []client.ResolveMatch, error) {
	return &client.ResolveResponse{}, nil, c.Failure
}
