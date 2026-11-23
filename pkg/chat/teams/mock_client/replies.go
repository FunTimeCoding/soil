package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) Replies(
	team string,
	channel string,
	identifier string,
	limit int,
) ([]*message.Message, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, m := range c.messages[key(team, channel)] {
		if m.Identifier == identifier {
			return limited(c.replies[key(team, channel, identifier)], limit), nil
		}
	}

	return nil, not_found.New("root message", identifier)
}
