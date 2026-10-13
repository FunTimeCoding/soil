package pool

func (p *Pool[T, P, C]) Set(
	instance string,
	c C,
) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.clients[instance] = c
}
