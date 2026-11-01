package held_peer

func (p *Peer) Refuse(e ...error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.refusal = append(p.refusal, e...)
}
