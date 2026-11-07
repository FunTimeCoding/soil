package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetPermissions(
	x context.Context,
	params *client.GetPermissionsParams,
) (*client.PermissionReport, error) {
	r, e := c.generated.GetPermissionsWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get permissions", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
