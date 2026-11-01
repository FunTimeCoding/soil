package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PostRelation(
	x context.Context,
	b client.PostRelationJSONRequestBody,
) error {
	r, e := c.client.PostRelationWithResponse(x, b)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusOK {
		return refusal.New("post relation", r.HTTPResponse, r.Body)
	}

	return nil
}
