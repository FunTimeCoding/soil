package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSessionDetail(
	x context.Context,
	identifier string,
) (*client.SessionDetailResponse, error) {
	r, e := c.generated.GetSessionDetailWithResponse(x, identifier)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get session detail", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
