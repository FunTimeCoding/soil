package runner

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/provision/constant"
)

func (r *Runner) syncTick() bool {
	changed := false
	v := recovery.Catch(func() { changed = r.gitSync() })

	if v == nil {
		if r.syncFailures > 0 {
			r.logger.Structured(
				"git_sync_recovered",
				constant.RunnerConsecutive,
				r.syncFailures,
			)
			r.syncFailures = 0
		}

		return changed
	}

	r.syncFailures++
	e, okay := v.(error)

	if !okay {
		e = fmt.Errorf("%v", v)
	}

	if r.syncFailures == 1 {
		r.reporter.CaptureException(e)
	} else {
		r.logger.Structured(
			"git_sync_failed",
			constant.RunnerError,
			e.Error(),
			constant.RunnerStderr,
			stderr(e),
			constant.RunnerConsecutive,
			r.syncFailures,
		)
	}

	r.recovery.Run(r.healRepository)

	return false
}
