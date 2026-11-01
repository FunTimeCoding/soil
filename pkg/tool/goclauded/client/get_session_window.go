package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSessionWindow(
	x context.Context,
	identifier string,
	params *client.GetSessionWindowParams,
) (*client.WindowResponse, error) {
	r, e := c.generated.GetSessionWindowWithResponse(x, identifier, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get session window", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
