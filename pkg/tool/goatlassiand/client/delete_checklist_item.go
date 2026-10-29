package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) DeleteChecklistItem(
	x context.Context,
	key string,
	index int,
) ([]*client.ChecklistItem, error) {
	r, e := c.client.DeleteChecklistItemWithResponse(x, key, index)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("delete checklist item", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
