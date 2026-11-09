package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) RestoreBackup(
	x context.Context,
	name string,
	body client.RestoreBackupJSONRequestBody,
) (*client.RestoreResult, error) {
	r, e := c.client.RestoreBackupWithResponse(x, name, body)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(constant.RestoreBackup, r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
