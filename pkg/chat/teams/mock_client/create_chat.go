package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/chat"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
)

func (c *Client) CreateChat(
	kind string,
	users []string,
	topic string,
) (*chat.Chat, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if kind == constant.TeamsDirectChat {
		for _, h := range c.chats {
			if h.Kind == constant.TeamsDirectChat &&
				c.membersAre(h.Identifier, users) {
				return h, nil
			}
		}
	}

	result := chat.New(
		fmt.Sprintf("19:created-%d@thread.v2", len(c.chats)+1),
		kind,
	)
	result.Topic = topic
	c.chats = append(c.chats, result)

	for _, u := range users {
		name := u

		if v, e := c.User(u); e == nil {
			name = v.Name
		} else if c.me != nil && c.me.Identifier == u {
			name = c.me.Name
		}

		c.chatMembers[result.Identifier] = append(
			c.chatMembers[result.Identifier],
			member.New(name, u),
		)
	}

	return result, nil
}
