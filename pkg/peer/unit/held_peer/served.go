package held_peer

func (p *Peer) Served() <-chan struct{} {
	return p.served
}
