package pending_digest

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"sync"
	"time"
)

type Pending struct {
	mutex    sync.Mutex
	events   map[uint][]*event.Event
	roots    map[uint]string
	next     map[uint]time.Time
	attempts map[uint]int
	base     time.Duration
	limit    time.Duration
}
