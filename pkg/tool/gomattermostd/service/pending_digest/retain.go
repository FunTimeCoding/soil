package pending_digest

import "slices"

func (p *Pending) Retain(
	root string,
	identifiers []uint,
) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for identifier, v := range p.roots {
		if v == root && !slices.Contains(identifiers, identifier) {
			delete(p.events, identifier)
			delete(p.roots, identifier)
			delete(p.next, identifier)
			delete(p.attempts, identifier)
		}
	}
}
