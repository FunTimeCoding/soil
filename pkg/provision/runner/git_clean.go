package runner

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func (r *Runner) gitClean() {
	lock := filepath.Join(r.cloneDirectory, constant.RunnerIndexLock)

	if !system.FileExists(lock) {
		return
	}

	system.RemoveFile(lock)
	r.logger.Structured("git_lock_removed", constant.RunnerPath, lock)
}
