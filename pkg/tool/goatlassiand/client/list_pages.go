package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListPages(
	x context.Context,
	space string,
	status string,
) ([]*client.ConfluencePage, error) {
	params := &client.ListPagesParams{Space: space}

	if status != "" {
		params.Status = &status
	}

	r, e := c.client.ListPagesWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("pages", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
