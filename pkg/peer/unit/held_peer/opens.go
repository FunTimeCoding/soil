package held_peer

func (p *Peer) Opens() int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.opened
}
