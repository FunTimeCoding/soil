package runner

import "github.com/funtimecoding/soil/pkg/provision/constant"

func (r *Runner) forceClone() {
	r.logger.Structured(
		"force_clone",
		constant.RunnerConsecutive,
		r.syncFailures,
	)
	r.installClone(r.stageClone())
	r.gitConfigure()
	r.syncFailures = 0
}
