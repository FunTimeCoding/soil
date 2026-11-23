package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) ChannelMessages(
	team string,
	channel string,
	limit int,
) ([]*message.Message, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if !c.hasChannel(team, channel) {
		return nil, not_found.New("channel", channel)
	}

	return limited(c.messages[key(team, channel)], limit), nil
}
