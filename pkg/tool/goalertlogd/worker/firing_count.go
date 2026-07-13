package worker

func (w *Worker) FiringCount() int {
	return w.store.MustUnresolvedCount()
}
