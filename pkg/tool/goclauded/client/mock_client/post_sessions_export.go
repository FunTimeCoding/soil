package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostSessionsExport(_ context.Context) (*client.ExportResponse, error) {
	return &client.ExportResponse{}, c.Failure
}
