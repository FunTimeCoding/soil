package telegram

import "github.com/funtimecoding/soil/pkg/chat/telegram/message"

func (c *Client) Messages() []*message.Message {
	var result []*message.Message

	for _, u := range c.Updates() {
		if u.Message != nil {
			m := message.New(u.Message)
			m.Update = &u
			result = append(result, m)

			if h := u.FromChat(); h != nil {
				c.saveChannel(h.ID, h.Title)
			}

			if s := u.SentFrom(); s != nil {
				c.saveUser(s.ID, s.UserName)
			}
		}
	}

	return result
}
