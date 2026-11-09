package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListBackups(
	x context.Context,
	params *client.ListBackupsParams,
) (*client.BackupList, error) {
	r, e := c.client.ListBackupsWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(constant.ListBackups, r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
