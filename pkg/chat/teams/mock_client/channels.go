package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/channel"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) Channels(team string) ([]*channel.Channel, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if !c.hasTeam(team) {
		return nil, not_found.New("team", team)
	}

	return append([]*channel.Channel{}, c.channels[team]...), nil
}
