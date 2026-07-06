package connection

import (
	"github.com/funtimecoding/soil/pkg/assistant/connection/subscribe_command"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Connection) Subscribe(
	event string,
	s Subscriber,
) {
	cmd := subscribe_command.New()
	cmd.Type = Subscribe
	cmd.Event = event
	_, e := c.send(cmd, s)
	errors.PanicOnError(e)
}
