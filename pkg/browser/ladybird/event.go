package ladybird

import "github.com/funtimecoding/soil/pkg/browser/types/message"

func (c *Client) event(r *message.Reply) {
	if c.option.Event != nil {
		c.option.Event(r)
	}
}
