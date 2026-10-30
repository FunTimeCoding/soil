package worker

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func (w *Worker) poll() {
	_, e := w.service.IndexCollections("")
	errors.PanicOnError(e)

	if pruned := w.service.Prune(); pruned > 0 {
		w.logger.Plain(constant.PrunedContent, pruned)
	}

	f := w.service.EmbedPending(w.context)

	switch {
	case f == nil, w.context.Err() != nil:
	case classify.Reportable(f):
		w.reporter.CaptureException(f)
	default:
		w.logger.Plain(constant.EmbeddingDeferred, f)
	}
}
