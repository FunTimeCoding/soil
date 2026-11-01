package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetCoverage(
	x context.Context,
) (*client.CoverageResponse, error) {
	r, e := c.generated.GetCoverageWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get coverage", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
