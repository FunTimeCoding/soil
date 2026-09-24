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
	next := join.Empty(r.clonePath, constant.RunnerStageSuffix)
	r.removeStage(next)
	system.CopyDirectory(stage, next)

	if system.DirectoryExists(r.clonePath) {
		r.quarantine()
	}

	errors.PanicOnError(os.Rename(next, r.clonePath))
}
