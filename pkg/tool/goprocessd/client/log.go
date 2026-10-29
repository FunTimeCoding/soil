package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Log(
	x context.Context,
	name string,
	all bool,
) (*client.Log, error) {
	params := &client.GetLogParams{}

	if all {
		params.All = &all
	}

	r, e := c.client.GetLogWithResponse(x, name, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("log", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
