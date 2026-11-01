package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostSessionExport(
	x context.Context,
	identifier string,
) (*client.ExportResponse, error) {
	r, e := c.generated.PostSessionExportWithResponse(x, identifier)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post session export", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
