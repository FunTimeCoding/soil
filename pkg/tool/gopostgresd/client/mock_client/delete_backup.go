package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
)

func (c *Client) DeleteBackup(
	_ context.Context,
	_ string,
	_ *client.DeleteBackupParams,
) error {
	return c.Failure
}
