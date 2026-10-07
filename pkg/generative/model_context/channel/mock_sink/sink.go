package mock_sink

import (
	"github.com/funtimecoding/soil/pkg/generative/types/sink_event"
	"sync"
)

type Sink struct {
	mutex sync.Mutex
	event []sink_event.Event
}
