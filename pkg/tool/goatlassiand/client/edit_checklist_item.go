package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) EditChecklistItem(
	x context.Context,
	key string,
	index int,
	text string,
) ([]*client.ChecklistItem, error) {
	r, e := c.client.EditChecklistItemWithResponse(
		x,
		key,
		index,
		client.ChecklistTextRequest{Text: text},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("edit checklist item", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
