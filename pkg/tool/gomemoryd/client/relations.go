package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *RestClient) Relations(x context.Context) ([]client.Relation, error) {
	r, e := c.http.GetRelationsWithResponse(x, &client.GetRelationsParams{})

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("relations", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
