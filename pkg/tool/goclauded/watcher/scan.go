package watcher

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/sweep"
)

func (w *Watcher) scan() {
	sweep.Run(w.harbor)
	w.service.PopulateCache()
	w.service.BackfillSessions()
	w.service.CheckConsistency()

	for _, e := range w.service.CatchUpSearch(context.Background()) {
		if classify.Reportable(e) {
			w.reporter.CaptureException(e)
		}
	}
}
