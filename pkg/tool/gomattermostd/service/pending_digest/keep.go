package pending_digest

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"time"
)

func (p *Pending) Keep(
	identifier uint,
	root string,
	events []*event.Event,
	now time.Time,
) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.attempts[identifier]++
	delay := p.base

	for range p.attempts[identifier] - 1 {
		if delay >= p.limit {
			break
		}

		delay *= 2
	}

	p.events[identifier] = events
	p.roots[identifier] = root
	p.next[identifier] = now.Add(min(delay, p.limit))
}
