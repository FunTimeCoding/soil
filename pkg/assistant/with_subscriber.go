package assistant

import "github.com/funtimecoding/soil/pkg/assistant/connection"

func WithSubscriber(h connection.Subscriber) Option {
	return func(c *Client) {
		c.subscriber = h
	}
}
