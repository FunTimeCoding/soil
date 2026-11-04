package pending_digest

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"

func (p *Pending) Take(identifier uint) []*event.Event {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	result := p.events[identifier]
	delete(p.events, identifier)
	delete(p.roots, identifier)
	delete(p.next, identifier)

	return result
}
