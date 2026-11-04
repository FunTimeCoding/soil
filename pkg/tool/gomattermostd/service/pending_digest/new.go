package pending_digest

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"time"
)

func New(
	base time.Duration,
	limit time.Duration,
) *Pending {
	return &Pending{
		events:   map[uint][]*event.Event{},
		roots:    map[uint]string{},
		next:     map[uint]time.Time{},
		attempts: map[uint]int{},
		base:     base,
		limit:    limit,
	}
}
