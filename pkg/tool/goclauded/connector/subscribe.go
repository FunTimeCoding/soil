package connector

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
)

func (c *Client) Subscribe(
	name string,
	kinds []string,
) (<-chan *stream_event.Event, func()) {
	events := make(chan *stream_event.Event, constant.StreamBuffer)
	x, cancel := context.WithCancel(context.Background())
	go func() {
		defer close(events)
		c.streamLoop(x, name, kinds, events)
	}()

	return events, cancel
}
