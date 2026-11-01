package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostSessionLabel(
	x context.Context,
	identifier string,
	body client.PostSessionLabelJSONRequestBody,
) (*client.LabelResponse, error) {
	r, e := c.generated.PostSessionLabelWithResponse(x, identifier, body)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post session label", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
