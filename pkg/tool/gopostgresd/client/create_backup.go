package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateBackup(
	x context.Context,
	body client.CreateBackupJSONRequestBody,
) (*client.Backup, error) {
	r, e := c.client.CreateBackupWithResponse(x, body)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(constant.CreateBackup, r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
