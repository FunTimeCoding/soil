package worker

func (w *Worker) Close() {
	w.once.Do(func() { w.Face.Close() })
}
