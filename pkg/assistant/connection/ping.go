package connection

import "github.com/funtimecoding/soil/pkg/assistant/connection/ping_command"

func (c *Connection) Ping() error {
	p := ping_command.New()
	p.Type = Ping
	_, e := c.send(p, nil)

	return e
}
