package assistant

import (
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
)

func (c *Client) deliver(m *message.Message) {
	if v := recovery.Catch(func() { c.subscriber(m) }); v != nil {
		panic(withEventContext(m, v))
	}
}
