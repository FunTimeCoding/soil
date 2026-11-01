package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetMessages(
	x context.Context,
	params *client.GetMessagesParams,
) (*client.MessageReadResponse, error) {
	r, e := c.generated.GetMessagesWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get messages", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
