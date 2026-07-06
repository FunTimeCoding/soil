package watcher

import "github.com/funtimecoding/soil/pkg/tool/goclauded/sweep"

func (w *Watcher) scan() {
	sweep.Run(w.harbor)
	w.service.PopulateCache()
	w.service.BackfillSessions()
	w.service.CheckConsistency()
}
