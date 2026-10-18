package unit

import "context"

func (p *SwitchablePeer) Reachable(context.Context) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.probes++

	return p.failure
}
