package worker

import "time"

func (w *Worker) Start() {
	go func() {
		t := time.NewTicker(w.interval)
		defer t.Stop()
		w.recovery.Run(w.Poll)

		for {
			select {
			case <-w.context.Done():
				return
			case <-t.C:
				w.recovery.Run(w.Poll)
			}
		}
	}()
}
