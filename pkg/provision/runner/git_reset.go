package runner

import "github.com/funtimecoding/soil/pkg/provision/constant"

func (r *Runner) gitReset() {
	c := r.newRun()
	c.Directory = r.cloneDirectory
	c.Start("git", "reset", "--hard", constant.RunnerRemoteBranch)
}
