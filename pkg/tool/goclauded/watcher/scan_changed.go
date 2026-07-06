package watcher

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/sweep"
)

func (w *Watcher) scanChanged() {
	for identifier, sourcePath := range w.changed {
		if sourcePath != "" {
			if sweep.SyncFile(sourcePath, w.harbor) {
				w.logger.Structured(
					"source_file_shrunk",
					constant.Identifier,
					identifier,
				)
			}
		}

		w.service.EnrichSession(identifier)
	}

	w.changed = make(map[string]string)
}
