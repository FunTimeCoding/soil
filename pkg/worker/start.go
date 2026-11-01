package worker

import "time"

func (w *Worker) Start() {
	w.started.Do(
		func() {
			go func() {
				defer close(w.finished)

				for {
					w.recovery.Run(w.Poll)

					if w.once {
						return
					}

					select {
					case <-w.context.Done():
						return
					case <-time.After(w.interval):
					}
				}
			}()
		},
	)
}
