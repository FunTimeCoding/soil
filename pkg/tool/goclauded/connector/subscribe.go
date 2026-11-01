package connector

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/hold"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector/stream"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
)

func (c *Client) Subscribe(
	name string,
	kinds []string,
	l *logger.Logger,
	r face.Reporter,
) (<-chan *stream_event.Event, func()) {
	events := make(chan *stream_event.Event, constant.StreamBuffer)
	h := hold.New(
		c.base,
		stream.New(c.pool, c.streamLocator(name, kinds), c.token, events),
		l,
		r,
	)
	h.Start()
	go func() {
		<-h.Done()
		close(events)
	}()

	return events, h.Close
}
