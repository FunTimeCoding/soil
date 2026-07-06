package connection

import "github.com/funtimecoding/soil/pkg/assistant/message"

func (c *Connection) consumeResult(m *message.Message) {
	c.RLock()
	defer c.RUnlock()

	if !m.Success && m.Error != nil {
		// command failed
		return
	}
}
