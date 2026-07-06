package telegram

import "github.com/funtimecoding/soil/pkg/chat/telegram/message"

func (c *Client) MessagesByChannel(name string) []*message.Message {
	var result []*message.Message

	for _, m := range c.Messages() {
		if m.Raw.Chat.Title != name {
			continue
		}

		if m.Text == "" {
			continue
		}

		result = append(result, m)
	}

	return result
}
