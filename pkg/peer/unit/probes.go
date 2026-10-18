package unit

func (p *SwitchablePeer) Probes() int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.probes
}
