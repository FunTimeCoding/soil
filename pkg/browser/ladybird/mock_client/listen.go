package mock_client

import "github.com/funtimecoding/soil/pkg/browser/types/message"

func (c *Client) Listen(event func(*message.Reply)) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.event = event
}
