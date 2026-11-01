package reconciler

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
)

func (r *Reconciler) fail(
	hypervisor string,
	e error,
) {
	if connection.Classify(e) != nil || unavailable.Is(e) ||
		!classify.Reportable(e) {
		r.logger.Plain("reconcile %s: %v", hypervisor, e)

		return
	}

	r.reporter.CaptureException(e)
}
