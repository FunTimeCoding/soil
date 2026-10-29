package worker

import "time"

func (w *Worker) Start() {
	go func() {
		t := time.NewTicker(w.interval)
		defer t.Stop()
		w.recovery.Run(w.Poll)

		for {
			select {
			case <-t.C:
				w.recovery.Run(w.Poll)
			case <-w.stop:
				return
			}
		}
	}()
}
