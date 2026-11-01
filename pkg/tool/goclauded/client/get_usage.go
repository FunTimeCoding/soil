package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) GetUsage(
	x context.Context,
) (*client.UsageResponse, bool, error) {
	r, e := c.generated.GetUsageWithResponse(x)

	if e != nil {
		return nil, false, e
	}

	if r.StatusCode() == http.StatusNoContent {
		return nil, false, nil
	}

	if r.JSON200 == nil {
		return nil, false, refusal.New("get usage", r.HTTPResponse, r.Body)
	}

	return r.JSON200, true, nil
}
