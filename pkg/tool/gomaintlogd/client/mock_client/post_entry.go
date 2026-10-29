package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
)

func (c *Client) PostEntry(
	_ context.Context,
	b client.PostEntryRequest,
) (*client.EntryResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.EntryResponse{Action: b.Action, User: b.User}, nil
}
