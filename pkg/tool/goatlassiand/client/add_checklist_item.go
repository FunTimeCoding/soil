package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) AddChecklistItem(
	x context.Context,
	key string,
	text string,
) ([]*client.ChecklistItem, error) {
	r, e := c.client.AddChecklistItemWithResponse(
		x,
		key,
		client.ChecklistTextRequest{Text: text},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("add checklist item", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
