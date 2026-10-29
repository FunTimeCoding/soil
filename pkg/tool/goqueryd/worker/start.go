package worker

import "time"

func (w *Worker) Start() {
	go func() {
		t := time.NewTicker(w.interval)
		defer t.Stop()
		w.recovery.Run(w.poll)

		for {
			select {
			case <-t.C:
				w.recovery.Run(w.poll)
			case <-w.stop:
				return
			}
		}
	}()
}
