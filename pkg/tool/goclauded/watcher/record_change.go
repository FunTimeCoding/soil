package watcher

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"path/filepath"
	"strings"
)

func (w *Watcher) recordChange(path string) {
	name := filepath.Base(path)

	if !strings.HasSuffix(name, constant.NotationLogExtension) {
		return
	}

	identifier := strings.TrimSuffix(name, constant.NotationLogExtension)

	if strings.HasPrefix(path, w.harbor) {
		if _, exists := w.changed[identifier]; !exists {
			w.changed[identifier] = ""
		}

		return
	}

	w.changed[identifier] = path
}
