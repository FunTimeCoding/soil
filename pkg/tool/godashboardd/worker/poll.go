package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
)

func (w *Worker) Poll() {
	for _, e := range w.service.Refresh() {
		if classify.Reportable(e) {
			w.reporter.CaptureException(e)

			continue
		}

		w.logger.Plain(constant.RefreshFailed, e)
	}
}
