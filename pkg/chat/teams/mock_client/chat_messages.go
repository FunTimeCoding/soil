package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) ChatMessages(
	chat string,
	limit int,
) ([]*message.Message, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, h := range c.chats {
		if h.Identifier == chat {
			return limited(c.chatMessages[chat], limit), nil
		}
	}

	return nil, not_found.New("chat", chat)
}
