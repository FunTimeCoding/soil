package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessionsSearch(_ context.Context, _ *client.GetSessionsSearchParams) (*client.SearchResponse, error) {
	return &client.SearchResponse{}, c.Failure
}
