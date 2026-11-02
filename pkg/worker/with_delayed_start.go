package worker

func (w *Worker) WithDelayedStart() *Worker {
	w.delayed = true

	return w
}
