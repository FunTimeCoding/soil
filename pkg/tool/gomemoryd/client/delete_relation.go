package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) DeleteRelation(
	x context.Context,
	p *client.DeleteRelationParams,
) error {
	r, e := c.client.DeleteRelationWithResponse(x, p)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusOK {
		return refusal.New("delete relation", r.HTTPResponse, r.Body)
	}

	return nil
}
