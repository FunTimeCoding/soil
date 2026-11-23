package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/chat"
)

func (c *Client) Chats(limit int) ([]*chat.Chat, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if limit <= 0 {
		limit = constant.TeamsDefaultLimit
	}

	return append([]*chat.Chat{}, c.chats[:min(limit, len(c.chats))]...), nil
}
