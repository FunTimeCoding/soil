package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
)

func (w *Worker) logFailure(
	hypervisor string,
	e error,
) {
	if connection.Classify(e) != nil || unavailable.Is(e) ||
		!classify.Reportable(e) {
		w.log.Plain("poll hypervisor %s: %v", hypervisor, e)

		return
	}

	w.reporter.CaptureException(e)
}
