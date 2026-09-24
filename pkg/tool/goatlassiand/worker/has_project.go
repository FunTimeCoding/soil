package worker

func (w *Worker) HasProject() bool {
	return len(w.project) > 0
}
