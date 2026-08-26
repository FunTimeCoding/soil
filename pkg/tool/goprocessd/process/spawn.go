package process

func (p *Process) Spawn(
	environment []string,
	onExit func(supervisorStopped bool),
) bool {
	p.mutex.Lock()

	if p.handle != nil {
		p.mutex.Unlock()

		return false
	}

	go func() {
		supervisorStopped := p.run(environment)
		p.mutex.Unlock()

		if onExit != nil {
			onExit(supervisorStopped)
		}
	}()

	return true
}
