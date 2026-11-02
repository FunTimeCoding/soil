package pool

func (p *Pool[T, P, C]) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, c := range p.clients {
		release(c)
	}

	clear(p.clients)
}
