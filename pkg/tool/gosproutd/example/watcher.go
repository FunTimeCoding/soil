package example

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
	"time"
)

func Watcher() {
	directory := filepath.Join(os.TempDir(), "sprout-watcher-test")
	system.MakeDirectory(directory)
	console.Format("watching: %s\n", directory)
	w, e := fsnotify.NewWatcher()
	errors.PanicOnError(e)
	defer func() { errors.PanicOnError(w.Close()) }()
	errors.PanicOnError(w.Add(directory))
	console.Format(
		"ready - create, modify, rename, or delete files in the directory\n\n",
	)

	for {
		select {
		case v, okay := <-w.Events:
			if !okay {
				return
			}

			console.Format(
				"%s  %-10s  %s\n",
				time.Now().Format("15:04:05.000"),
				v.Op.String(),
				filepath.Base(v.Name),
			)
		case f, okay := <-w.Errors:
			if !okay {
				return
			}

			console.Format(
				"%s  ERROR      %v\n",
				time.Now().Format("15:04:05.000"),
				f,
			)
		}
	}
}
