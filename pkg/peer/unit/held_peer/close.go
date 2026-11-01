package held_peer

func (p *Peer) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.closed++
}
