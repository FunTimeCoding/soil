package held_peer

func (p *Peer) Closes() int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.closed
}
