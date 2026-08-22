package runner

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

// The runner is the only git user in its container, so a lock
// present between commands survived a killed git process.
func (r *Runner) gitClean() {
	lock := filepath.Join(r.clonePath, constant.RunnerIndexLock)

	if !system.FileExists(lock) {
		return
	}

	system.RemoveFile(lock)
	r.logger.Structured("git_lock_removed", constant.RunnerPath, lock)
}
