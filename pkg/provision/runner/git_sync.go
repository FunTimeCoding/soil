package runner

import "github.com/funtimecoding/soil/pkg/provision/constant"

func (r *Runner) gitSync() bool {
	r.gitFetch()
	local := r.gitRevision("HEAD")
	remote := r.gitRevision(constant.RunnerRemoteBranch)

	if local == remote {
		r.logger.Structured("git_sync", constant.RunnerStatus, "unchanged")

		return false
	}

	r.logger.Structured(
		"git_sync",
		constant.RunnerStatus,
		"changed",
		constant.RunnerLocal,
		local,
		constant.RunnerRemote,
		remote,
	)
	r.gitPull()

	return true
}
