package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func (w *Worker) poll() {
	pruned, e := w.maintain(w.context)
	w.metric.CountPruned(pruned)

	switch {
	case e == nil, w.context.Err() != nil:
	case classify.Reportable(e):
		w.reporter.CaptureException(e)
	default:
		w.logger.Plain(constant.EmbeddingDeferred, e)
	}
}
