package runner

import "github.com/funtimecoding/soil/pkg/provision/constant"

func (r *Runner) syncWithDiff() *SyncResult {
	r.gitClean()
	r.gitFetch()
	local := r.gitRevision("HEAD")
	remote := r.gitRevision(constant.RunnerRemoteBranch)

	if local == remote {
		r.logger.Structured("sync", constant.RunnerStatus, "unchanged")

		return &SyncResult{}
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

	return &SyncResult{Changed: true, Diff: diff}
}
