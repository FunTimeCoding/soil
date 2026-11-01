package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostSessionExport(_ context.Context, _ string) (*client.ExportResponse, error) {
	return &client.ExportResponse{}, c.Failure
}
