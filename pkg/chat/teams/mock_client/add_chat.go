package mock_client

import "github.com/funtimecoding/soil/pkg/chat/teams/chat"

func (c *Client) AddChat(h *chat.Chat) {
	c.chats = append(c.chats, h)
}
