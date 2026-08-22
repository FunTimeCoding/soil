package runner

import "github.com/funtimecoding/soil/pkg/system/run"

func (r *Runner) newRun() *run.Run {
	result := run.New()

	if r.registry != nil {
		result.WithRegistry(r.registry)
	}

	return result
}
