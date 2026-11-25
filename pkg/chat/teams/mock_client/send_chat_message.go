package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) SendChatMessage(
	chat string,
	text string,
) (*message.Message, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, h := range c.chats {
		if h.Identifier == chat {
			result := message.New(
				fmt.Sprintf("sent-%d", len(c.chatMessages[chat])+1),
				"",
				text,
			)
			c.chatMessages[chat] = append(c.chatMessages[chat], result)

			return result, nil
		}
	}

	return nil, not_found.New("chat", chat)
}
