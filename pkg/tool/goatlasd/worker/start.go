package worker

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"time"
)

func (w *Worker) Start() {
	w.logger.Plain(constant.StartMessage)
	go func() {
		t := time.NewTicker(w.interval)
		defer t.Stop()
		w.recovery.Run(w.poll)

		for {
			select {
			case <-w.context.Done():
				return
			case <-t.C:
				w.recovery.Run(w.poll)
			}
		}
	}()
}
