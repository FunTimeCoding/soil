package switchable_peer

func (p *Peer) Fail(e error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.failure = e
}
