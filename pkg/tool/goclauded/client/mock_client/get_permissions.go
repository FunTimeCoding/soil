package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetPermissions(_ context.Context, _ *client.GetPermissionsParams) (*client.PermissionReport, error) {
	return &client.PermissionReport{}, c.Failure
}
