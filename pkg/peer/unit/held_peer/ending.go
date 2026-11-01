package held_peer

func (p *Peer) ending() chan error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.end == nil {
		p.end = make(chan error, 1)
	}

	return p.end
}
