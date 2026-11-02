package pool

func (p *Pool[T, P, C]) Forget(
	instance string,
	c C,
) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if held, okay := p.clients[instance]; !okay || held != c {
		return
	}

	release(c)
	delete(p.clients, instance)
}
