package held_peer

import "context"

func (p *Peer) Open(context.Context) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.opened++

	if len(p.refusal) == 0 {
		return nil
	}

	e := p.refusal[0]
	p.refusal = p.refusal[1:]

	return e
}
