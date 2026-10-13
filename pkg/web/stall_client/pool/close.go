package pool

func (p *Pool) Close() {
	p.transport.CloseIdleConnections()
}
