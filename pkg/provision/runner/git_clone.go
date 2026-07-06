package runner

import "os"

func (r *Runner) gitClone() {
	if _, e := os.Stat(r.clonePath); e == nil {
		r.gitPull()

		return
	}

	r.logger.Structured("git_clone", "repository", r.repository)
	r.newRun().Start("git", "clone", r.repository, r.clonePath)
}
