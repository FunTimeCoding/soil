package runner

import "github.com/funtimecoding/soil/pkg/system"

func (r *Runner) gitClone() {
	if system.DirectoryExists(r.clonePath) && r.validRepository() {
		r.gitConfigure()
		r.gitClean()
		r.gitFetch()
		r.gitReset()

		return
	}

	r.logger.Structured("git_clone", "repository", r.repository)
	r.installClone(r.stageClone())
	r.gitConfigure()
}
