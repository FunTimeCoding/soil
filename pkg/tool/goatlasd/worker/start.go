package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"time"
)

func (w *Worker) Start() {
	w.logger.Plain(constant.StartMessage)
	x, cancel := context.WithCancel(context.Background())
	w.cancel = cancel
	go func() {
		t := time.NewTicker(w.interval)
		defer t.Stop()
		w.recovery.Run(w.poll)

		for {
			select {
			case <-x.Done():
				return
			case <-t.C:
				w.recovery.Run(w.poll)
			}
		}
	}()
}
