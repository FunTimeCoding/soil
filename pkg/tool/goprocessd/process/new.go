package process

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/log"
	"os"
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
		logger:     log.New(os.Stdout, name, colorIndex, maxNameWidth),
	}
	p.condition = sync.NewCond(&p.mutex)

	return p
}
