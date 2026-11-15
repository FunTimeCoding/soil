package switchable_peer

func (p *Peer) Probes() int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.probes
}
