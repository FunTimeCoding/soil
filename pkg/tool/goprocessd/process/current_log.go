package process

func (p *Process) CurrentLog() ([]string, int) {
	return p.logger.Current()
}
