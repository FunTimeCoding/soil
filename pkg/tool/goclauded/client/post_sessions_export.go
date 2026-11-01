package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostSessionsExport(
	x context.Context,
) (*client.ExportResponse, error) {
	r, e := c.generated.PostSessionsExportWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post sessions export", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
