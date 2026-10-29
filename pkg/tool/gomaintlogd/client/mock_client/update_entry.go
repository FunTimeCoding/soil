package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
)

func (c *Client) UpdateEntry(
	_ context.Context,
	identifier int,
	b client.PostEntryRequest,
) (*client.EntryResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.EntryResponse{
		Id:     identifier,
		Action: b.Action,
		User:   b.User,
	}, nil
}
