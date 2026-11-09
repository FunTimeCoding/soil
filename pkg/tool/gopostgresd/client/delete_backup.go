package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) DeleteBackup(
	x context.Context,
	name string,
	params *client.DeleteBackupParams,
) error {
	r, e := c.client.DeleteBackupWithResponse(x, name, params)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New(constant.DeleteBackup, r.HTTPResponse, r.Body)
	}

	return nil
}
