package stream

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"io"
	"sync"
)

type Stream struct {
	mutex   sync.Mutex
	pool    *pool.Pool
	locator string
	token   string
	events  chan<- *stream_event.Event
	last    string
	body    io.ReadCloser
}
