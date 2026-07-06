package process

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/log"
	"sync"
)

func New(
	name string,
	command string,
	colorIndex int,
	maxNameWidth int,
) *Process {
	p := &Process{
		Name:       name,
		Command:    command,
		ColorIndex: colorIndex,
		logger:     log.New(name, colorIndex, maxNameWidth),
	}
	p.condition = sync.NewCond(&p.mutex)

	return p
}
