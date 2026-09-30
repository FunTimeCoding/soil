package assistant

import "github.com/funtimecoding/soil/pkg/assistant/message"

func (c *Client) dispatch(m *message.Message) {
	c.recovery.Run(func() { c.deliver(m) })
}
