package connection

import (
	"github.com/funtimecoding/soil/pkg/assistant/connection/subscribe_command"
	"github.com/funtimecoding/soil/pkg/assistant/constant"
)

func (c *Connection) Subscribe(
	event string,
	s Subscriber,
) error {
	cmd := subscribe_command.New()
	cmd.Type = constant.Subscribe
	cmd.Event = event
	_, e := c.send(cmd, s)

	return e
}
