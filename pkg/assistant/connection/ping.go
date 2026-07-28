package connection

import (
	"github.com/funtimecoding/soil/pkg/assistant/connection/ping_command"
	"github.com/funtimecoding/soil/pkg/assistant/constant"
)

func (c *Connection) Ping() error {
	p := ping_command.New()
	p.Type = constant.Ping
	_, e := c.send(p, nil)

	return e
}
