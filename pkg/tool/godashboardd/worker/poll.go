package worker

func (w *Worker) Poll() {
	w.service.Refresh()
}
