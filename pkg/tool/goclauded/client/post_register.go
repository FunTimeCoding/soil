package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostRegister(
	x context.Context,
	body client.PostRegisterJSONRequestBody,
) (*client.RegisterResponse, error) {
	r, e := c.generated.PostRegisterWithResponse(x, body)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post register", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
