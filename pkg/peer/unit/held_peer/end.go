package held_peer

func (p *Peer) End(e error) {
	p.end <- e
}
