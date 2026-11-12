package claim

import (
	"context"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim/message"
)

func (c *Connection) deliver(m message.Message) {
	c.mutex.Lock()
	x := c.context
	c.mutex.Unlock()

	if x == nil {
		x = context.Background()
	}

	select {
	case c.updates <- m:
	case <-x.Done():
	}
}
