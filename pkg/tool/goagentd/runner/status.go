package runner

import "github.com/funtimecoding/soil/pkg/tool/goagentd/types/run_result"

func (r *Runner) Status() *run_result.Result {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return run_result.New(r.state, r.result)
}
