package runner

import "github.com/funtimecoding/soil/pkg/tool/goagentd/constant"

func (r *Runner) Close() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.closed = true

	if r.state != constant.RunnerRunning {
		r.end()
	}
}
