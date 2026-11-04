package pending_digest

func (p *Pending) Done(identifier uint) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.attempts, identifier)
}
