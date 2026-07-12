package watcher

import (
	"github.com/funtimecoding/soil/pkg/git"
	"time"
)

func (w *Watcher) gitTimes() (map[string]time.Time, map[string]bool) {
	times, e := git.CommitTimes(w.seedDirectory)

	if e != nil {
		return nil, nil
	}

	uncommitted, f := git.UncommittedFiles(w.seedDirectory)

	if f != nil {
		return nil, nil
	}

	dirty := make(map[string]bool)

	for _, p := range uncommitted {
		dirty[p] = true
	}

	return times, dirty
}
