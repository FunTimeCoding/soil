package stream

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
)

func New(
	p *pool.Pool,
	locator string,
	token string,
	events chan<- *stream_event.Event,
) *Stream {
	return &Stream{pool: p, locator: locator, token: token, events: events}
}
