package assistant

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assistant/constant"
	"github.com/funtimecoding/soil/pkg/assistant/message"
)

func (c *Client) Start() {
	c.connection.Connect()
	h := c.subscriber
	wrap := h

	if c.reporter != nil {
		wrap = func(m *message.Message) {
			defer func() {
				if v := recover(); v != nil {
					c.reporter.CaptureWithContext(
						fmt.Errorf("%v", v),
						constant.EventContext,
						map[string]any{
							constant.TypeKey: m.Event.Type,
							constant.RawKey:  string(m.Event.Raw),
						},
					)
				}
			}()
			h(m)
		}
	}

	c.connection.Subscribe("", wrap)
	go c.connection.Read()
	go c.poll()
}
