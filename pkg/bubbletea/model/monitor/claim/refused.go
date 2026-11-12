package claim

import "github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim/message"

func (c *Connection) Refused(e error) {
	c.deliver(message.Message{Error: e})
}
