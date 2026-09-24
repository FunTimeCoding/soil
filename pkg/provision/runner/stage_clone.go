package runner

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"os"
)

func (r *Runner) stageClone() string {
	stage, e := os.MkdirTemp("", constant.RunnerStagePattern)
	errors.PanicOnError(e)
	r.logger.Structured("stage_clone", constant.RunnerPath, stage)
	c := r.newRun().NoPanic()
	c.Start("git", "clone", r.repository, stage)

	if c.Error != nil {
		r.removeStage(stage)
		errors.PanicOnError(c.Error)
	}

	return stage
}
