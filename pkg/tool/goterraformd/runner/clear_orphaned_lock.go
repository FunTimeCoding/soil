package runner

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/types/lock_detail"
	"path/filepath"
)

func (r *Runner) clearOrphanedLock() {
	l := r.stateLease()

	if l == nil || !l.Held() {
		return
	}

	d := lock_detail.New(l)

	if d == nil || !r.orphaned(d) {
		return
	}

	c := r.newRun().NoPanic()
	c.Directory = filepath.Join(r.clonePath, r.terraformPath)
	c.Start(
		r.executable,
		constant.ForceUnlock,
		constant.ForceArgument,
		d.Identifier,
	)
	context := map[string]any{
		"lock":      d.Identifier,
		"who":       d.Who,
		"operation": d.Operation,
		"created":   d.Created,
	}

	if c.Error != nil {
		r.reporter.CaptureWithContext(c.Error, constant.LockContext, context)

		return
	}

	r.metrics.LockCleared.Inc()
	r.reporter.CaptureWithContext(
		fmt.Errorf("terraform state lock orphaned, force-unlocked"),
		constant.LockContext,
		context,
	)
}
