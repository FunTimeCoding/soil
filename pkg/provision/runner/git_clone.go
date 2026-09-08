package runner

import "os"

func (r *Runner) gitClone() {
	if _, e := os.Stat(r.clonePath); e == nil {
		if r.validRepository() {
			r.gitConfigure()
			r.gitClean()
			r.gitFetch()
			r.gitReset()

			return
		}

		r.quarantine()
	}

	r.logger.Structured("git_clone", "repository", r.repository)
	r.newRun().Start("git", "clone", r.repository, r.clonePath)
	r.gitConfigure()
}
