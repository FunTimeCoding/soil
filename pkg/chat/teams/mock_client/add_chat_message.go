package mock_client

import "github.com/funtimecoding/soil/pkg/chat/teams/message"

func (c *Client) AddChatMessage(
	chat string,
	m *message.Message,
) {
	c.chatMessages[chat] = append(c.chatMessages[chat], m)
}
