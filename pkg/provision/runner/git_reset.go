package runner

import "github.com/funtimecoding/soil/pkg/provision/constant"

func (r *Runner) gitReset() {
	c := r.newRun()
	c.Directory = r.clonePath
	c.Start("git", "reset", "--hard", constant.RunnerRemoteBranch)
}
