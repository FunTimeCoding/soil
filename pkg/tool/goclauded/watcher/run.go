package watcher

import (
	"github.com/fsnotify/fsnotify"
	"time"
)

func (w *Watcher) run() {
	defer close(w.done)
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()
	debounce := time.NewTimer(0)
	<-debounce.C

	for {
		select {
		case <-w.context.Done():
			return
		case <-ticker.C:
			w.recovery.Run(
				func() {
					w.addDirectories()
					w.scan()
				},
			)
		case v, okay := <-w.notifier.Events:
			if !okay {
				return
			}

			if !v.Has(fsnotify.Create) && !v.Has(fsnotify.Write) {
				continue
			}

			w.recordChange(v.Name)
			debounce.Reset(1 * time.Second)
		case <-debounce.C:
			if len(w.changed) > 0 {
				w.recovery.Run(w.scanChanged)
			}
		case f, okay := <-w.notifier.Errors:
			if !okay {
				return
			}

			w.reporter.CaptureException(f)
		}
	}
}
