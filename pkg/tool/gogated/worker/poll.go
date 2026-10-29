package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
)

func (w *Worker) Poll() {
	for _, e := range w.service.Clean() {
		if classify.Reportable(e) {
			w.reporter.CaptureException(e)

			continue
		}

		w.logger.Plain(constant.CleanFailed, e)
	}
}
