package runner

import "github.com/funtimecoding/soil/pkg/provision/constant"

func (r *Runner) healRepository() {
	if r.syncFailures >= constant.RunnerHealThreshold {
		r.forceClone()

		return
	}

	if r.validRepository() {
		return
	}

	r.gitClone()
}
