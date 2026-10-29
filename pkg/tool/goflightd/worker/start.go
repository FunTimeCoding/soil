package worker

import "time"

func (w *Worker) Start() {
	go func() {
		for {
			w.recovery.Run(w.Poll)

			select {
			case <-w.context.Done():
				return
			case <-time.After(w.interval):
			}
		}
	}()
}
