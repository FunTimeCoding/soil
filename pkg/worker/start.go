package worker

func (w *Worker) Start() {
	w.started.Do(
		func() {
			go func() {
				defer close(w.finished)

				if w.at == nil && !w.delayed {
					w.recovery.Run(w.Poll)
				}

				for !w.once && w.sleep() {
					w.recovery.Run(w.Poll)
				}
			}()
		},
	)
}
