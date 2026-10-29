package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetChecklist(
	x context.Context,
	key string,
) ([]*client.ChecklistItem, error) {
	r, e := c.client.GetChecklistWithResponse(x, key)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("checklist", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
