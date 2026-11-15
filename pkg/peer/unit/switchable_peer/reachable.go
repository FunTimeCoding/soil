package switchable_peer

import "context"

func (p *Peer) Reachable(context.Context) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.probes++

	return p.failure
}
