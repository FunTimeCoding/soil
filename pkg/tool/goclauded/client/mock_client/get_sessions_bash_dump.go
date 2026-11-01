package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessionsBashDump(_ context.Context) (*client.BashDumpResponse, error) {
	return &client.BashDumpResponse{}, c.Failure
}
