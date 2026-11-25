package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetRoster(
	_ context.Context,
	_ *client.GetRosterParams,
) (*client.RosterResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.RosterResponse{
		Entries: append([]client.RosterEntry{}, c.Roster...),
	}, nil
}
