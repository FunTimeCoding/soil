package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
)

func (c *Client) RestoreBackup(
	_ context.Context,
	_ string,
	_ client.RestoreBackupJSONRequestBody,
) (*client.RestoreResult, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.RestoreResult{}, nil
}
