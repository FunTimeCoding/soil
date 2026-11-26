package mock_client

import "github.com/funtimecoding/soil/pkg/browser/types/message"

func (c *Client) Emit(r *message.Reply) {
	c.mutex.Lock()
	listen := c.event
	c.mutex.Unlock()

	if listen != nil {
		listen(r)
	}
}
