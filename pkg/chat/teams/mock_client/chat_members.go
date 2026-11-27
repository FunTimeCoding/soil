package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) ChatMembers(chat string) ([]*member.Member, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, h := range c.chats {
		if h.Identifier == chat {
			return c.chatMembers[chat], nil
		}
	}

	return nil, not_found.New("chat", chat)
}
