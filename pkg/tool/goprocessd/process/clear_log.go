package process

func (p *Process) ClearLog() {
	p.logger.Clear()
}
