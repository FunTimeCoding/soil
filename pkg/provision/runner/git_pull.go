package runner

import "github.com/funtimecoding/soil/pkg/provision/constant"

func (r *Runner) gitPull() {
	c := r.newRun()
	c.Directory = r.clonePath
	c.Start("git", "pull", "origin", constant.RunnerBranch)
}
