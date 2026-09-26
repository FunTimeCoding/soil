package runner

import "github.com/funtimecoding/soil/pkg/tool/goagentd/constant"

func (r *Runner) Start(intent string) bool {
	r.mutex.Lock()

	if r.state == constant.RunnerRunning {
		r.mutex.Unlock()

		return false
	}

	r.state = constant.RunnerRunning
	r.result = ""
	r.mutex.Unlock()
	go r.execute(intent)

	return true
}
