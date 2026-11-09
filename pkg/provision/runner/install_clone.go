package runner

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
)

func (r *Runner) installClone(stage string) {
	defer r.removeStage(stage)
	next := join.Empty(r.cloneDirectory, constant.RunnerStageSuffix)
	r.removeStage(next)
	system.CopyDirectory(stage, next)

	if r.copyModes {
		copyModes(stage, next)
	}

	if system.DirectoryExists(r.cloneDirectory) {
		r.quarantine()
	}

	errors.PanicOnError(os.Rename(next, r.cloneDirectory))
}
