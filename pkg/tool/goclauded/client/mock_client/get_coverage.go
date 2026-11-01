package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetCoverage(_ context.Context) (*client.CoverageResponse, error) {
	return &client.CoverageResponse{}, c.Failure
}
