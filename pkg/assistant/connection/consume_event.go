package connection

import "github.com/funtimecoding/soil/pkg/assistant/message"

func (c *Connection) consumeEvent(m *message.Message) {
	c.RLock()
	f, okay := c.subscribers[m.Identifier]
	c.RUnlock()

	if !okay {
		return
	}

	go f(m)
}
