package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
)

func (r *Runner) Begin(intent string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.closed {
		return unavailable.Format(constant.ShuttingDown)
	}

	if r.state == constant.RunnerRunning {
		return conflict.Format(constant.AlreadyRunning)
	}

	r.state = constant.RunnerRunning
	r.result = ""
	go r.recovery.Run(func() { r.execute(intent) })

	return nil
}
