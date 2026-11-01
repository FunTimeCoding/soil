package connection

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assistant/constant"
	"time"
)

func (c *Connection) keepAlive(x context.Context) {
	t := time.NewTicker(constant.PingInterval)
	defer t.Stop()

	for {
		select {
		case <-x.Done():
			return
		case <-t.C:
			if c.Ping() != nil {
				c.Close()

				return
			}
		}
	}
}
