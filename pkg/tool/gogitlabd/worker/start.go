package worker

import "time"

func (w *Worker) Start() {
	go func() {
		t := time.NewTicker(w.interval)
		defer t.Stop()
		w.recovery.Run(w.Poll)
		t.Reset(w.nextInterval())

		for {
			select {
			case <-t.C:
				w.recovery.Run(w.Poll)
				t.Reset(w.nextInterval())
			case <-w.stop:
				return
			}
		}
	}()
}
