package connection

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assistant/constant"
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
)

func (c *Connection) Serve(x context.Context) error {
	c.RLock()
	s := c.connection
	c.RUnlock()

	if s == nil {
		return unreachable.Format("connection not open")
	}

	y, stop := context.WithCancel(x)
	defer stop()
	go c.keepAlive(y)

	for {
		var m message.Message

		if e := s.ReadJSON(&m); e != nil {
			return e
		}

		switch m.Type {
		case constant.Result:
			c.consumeResult(&m)
		case constant.Event:
			c.consumeEvent(&m)
		}
	}
}
