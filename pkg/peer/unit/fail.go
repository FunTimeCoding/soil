package unit

func (p *SwitchablePeer) Fail(e error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.failure = e
}
