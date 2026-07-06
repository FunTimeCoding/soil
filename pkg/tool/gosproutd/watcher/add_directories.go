package watcher

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
)

func (w *Watcher) addDirectories(n *fsnotify.Watcher) {
	errors.PanicOnError(
		filepath.Walk(
			w.seedDirectory,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				if e != nil {
					return nil
				}

				if i.IsDir() {
					errors.PanicOnError(n.Add(path))
				}

				return nil
			},
		),
	)
}
