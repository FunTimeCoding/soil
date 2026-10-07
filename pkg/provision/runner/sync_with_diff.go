package runner

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/update"
)

func (r *Runner) syncWithDiff() *update.Result {
	r.gitClean()
	r.gitFetch()
	local := r.gitRevision("HEAD")
	remote := r.gitRevision(constant.RunnerRemoteBranch)

	if local == remote {
		r.logger.Structured("sync", constant.RunnerStatus, "unchanged")

		return update.NewResult()
	}

	r.logger.Structured(
		"sync",
		constant.RunnerStatus,
		"changed",
		constant.RunnerLocal,
		local,
		constant.RunnerRemote,
		remote,
	)
	diff := r.gitDiffLog(local, remote)
	r.gitReset()
	u := update.NewResult()
	u.Changed = true
	u.Diff = diff

	return u
}
