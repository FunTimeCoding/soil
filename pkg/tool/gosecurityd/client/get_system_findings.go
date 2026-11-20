package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSystemFindings(
	x context.Context,
	name string,
) (*client.SystemDetail, error) {
	r, e := c.client.GetSystemFindingsWithResponse(x, name)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("system findings", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
