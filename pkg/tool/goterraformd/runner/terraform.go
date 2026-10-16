package runner

import "github.com/funtimecoding/soil/pkg/system/run"

func (r *Runner) terraform(
	directory string,
	arguments ...string,
) *run.Run {
	result := r.newRun().NoPanic()
	result.Directory = directory
	result.Start(arguments...)

	return result
}
