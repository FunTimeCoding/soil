package runner

import "github.com/funtimecoding/soil/pkg/console/description"

func (r *Runner) Meta() *description.Description {
	return description.New("Runner", "Runner")
}
