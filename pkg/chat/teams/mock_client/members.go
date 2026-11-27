package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/destination"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
)

func (c *Client) Members(d *destination.Destination) ([]*member.Member, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if _, _, e := c.target(d); e != nil {
		return nil, e
	}

	if d.Chat != "" {
		return c.chatMembers[d.Chat], nil
	}

	return c.chatMembers[key(d.Team, d.Channel)], nil
}
