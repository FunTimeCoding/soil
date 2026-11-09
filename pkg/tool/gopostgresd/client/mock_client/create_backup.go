package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
)

func (c *Client) CreateBackup(
	_ context.Context,
	_ client.CreateBackupJSONRequestBody,
) (*client.Backup, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if len(c.Backups) == 0 {
		return &client.Backup{}, nil
	}

	return &c.Backups[0], nil
}
