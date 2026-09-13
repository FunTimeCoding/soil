package worker

func (w *Worker) notify() {
	if w.notifier == nil {
		return
	}

	w.notifier.Notify()
}
