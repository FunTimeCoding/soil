package memory

import (
	"github.com/funtimecoding/soil/pkg/errors/types/reported_event"
	"sync"
)

type Memory struct {
	mutex  sync.Mutex
	events []*reported_event.Event
}
