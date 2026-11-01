package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostBackfill(_ context.Context, _ *client.PostBackfillParams) (*client.BackfillResponse, error) {
	return &client.BackfillResponse{}, c.Failure
}
