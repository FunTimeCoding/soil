package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *RestClient) Reachable(x context.Context) error {
	r, e := c.http.GetVersionsWithResponse(
		x,
		&client.GetVersionsParams{Since: "2099-01-01T00:00:00Z"},
	)

	if e != nil {
		return e
	}

	if r.JSON200 == nil {
		return refusal.New("versions", r.HTTPResponse, r.Body)
	}

	return nil
}
