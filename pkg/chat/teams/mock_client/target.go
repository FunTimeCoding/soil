package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/destination"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) target(
	d *destination.Destination,
) (map[string][]*message.Message, string, error) {
	if d.Chat != "" {
		for _, h := range c.chats {
			if h.Identifier == d.Chat {
				return c.chatMessages, d.Chat, nil
			}
		}

		return nil, "", not_found.New("chat", d.Chat)
	}

	if !c.hasChannel(d.Team, d.Channel) {
		return nil, "", not_found.New("channel", d.Channel)
	}

	if d.Thread == "" {
		return c.messages, key(d.Team, d.Channel), nil
	}

	for _, m := range c.messages[key(d.Team, d.Channel)] {
		if m.Identifier == d.Thread {
			return c.replies, key(d.Team, d.Channel, d.Thread), nil
		}
	}

	return nil, "", not_found.New("root message", d.Thread)
}
