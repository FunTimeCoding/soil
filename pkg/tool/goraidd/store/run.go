package store

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"path/filepath"
	"strings"
	"time"
)

func (s *Store) run() {
	s.enrichExisting()
	s.cleanup()
	w, e := fsnotify.NewWatcher()
	errors.PanicOnError(e)
	defer func() { errors.PanicOnError(w.Close()) }()
	errors.PanicOnError(w.Add(s.logCachePath))
	errors.PanicOnError(w.Add(s.elitePath))
	s.logger.Structured(
		"watcher_start",
		"log_cache",
		s.logCachePath,
		"elite_insights",
		s.elitePath,
	)
	cleanup := time.NewTicker(24 * time.Hour)
	defer cleanup.Stop()

	for {
		select {
		case <-s.stop:
			return
		case <-cleanup.C:
			s.recovery.Run(s.cleanup)
		case v, okay := <-w.Events:
			if !okay {
				return
			}

			if !v.Has(fsnotify.Create) && !v.Has(fsnotify.Write) {
				continue
			}

			s.recovery.Run(
				func() {
					if filepath.Base(v.Name) == constant.LogFile {
						s.logger.Structured("log_cache_sync")
						s.syncLogCache()
					}

					if strings.HasSuffix(
						v.Name,
						constant.DetailedWvWKillSuffix,
					) {
						s.enrichFile(
							filepath.Dir(v.Name),
							filepath.Base(v.Name),
						)
					}
				},
			)
		case f, okay := <-w.Errors:
			if !okay {
				return
			}

			panic(f)
		}
	}
}
