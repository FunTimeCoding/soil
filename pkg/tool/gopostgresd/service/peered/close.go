package peered

func (p *Peered) Close() {
	p.guard.Close()
	p.client.Close()
}
