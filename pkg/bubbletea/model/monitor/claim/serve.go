package claim

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
)

func (c *Connection) Serve(x context.Context) error {
	c.mutex.Lock()
	b := c.body
	c.mutex.Unlock()

	if b == nil {
		return unreachable.Format("claims not open")
	}

	return client.ReadClaims(
		b,
		func(v []generated.Claim) {
			select {
			case c.updates <- Message{Claims: v}:
			case <-x.Done():
			}
		},
	)
}
