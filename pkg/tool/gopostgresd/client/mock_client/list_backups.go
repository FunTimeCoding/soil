package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
)

func (c *Client) ListBackups(
	_ context.Context,
	_ *client.ListBackupsParams,
) (*client.BackupList, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.BackupList{Backups: c.Backups}, nil
}
