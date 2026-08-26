package process

import "sync"

func Replace(previous *Process, command string) *Process {
	p := &Process{
		Name:       previous.Name,
		Command:    command,
		ColorIndex: previous.ColorIndex,
		logger:     previous.logger,
	}
	p.condition = sync.NewCond(&p.mutex)

	return p
}
