package worker

func (w *Worker) HasProject() bool {
	return w.service.HasProject()
}
