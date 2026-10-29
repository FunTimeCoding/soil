package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostDocument(
	x context.Context,
	b client.PostDocumentJSONRequestBody,
) (*client.PushOutcome, error) {
	r, e := c.client.PostDocumentWithResponse(x, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("push document", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
