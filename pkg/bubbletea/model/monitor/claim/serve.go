package claim

import (
	"context"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim/message"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	gomonitord "github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
)

func (c *Connection) Serve(x context.Context) error {
	c.mutex.Lock()
	b := c.body
	c.mutex.Unlock()

	if b == nil {
		return unreachable.Format("claims not open")
	}

	return gomonitord.ReadClaims(
		b,
		func(v []client.Claim) {
			select {
			case c.updates <- message.Message{Claims: v}:
			case <-x.Done():
			}
		},
	)
}
