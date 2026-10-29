package worker

func (w *Worker) Start() {
	w.running.Add(1)
	go func() {
		defer w.running.Done()
		w.recovery.Run(w.run)
	}()
}
