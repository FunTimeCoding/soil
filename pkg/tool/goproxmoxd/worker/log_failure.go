package worker

func (w *Worker) logFailure(
	hypervisor string,
	e error,
) {
	if edge := w.reacher.Observe(hypervisor, e); edge != nil {
		w.log.Plain("poll hypervisor %s", edge)

		return
	}

	if down, _ := w.reacher.Down(hypervisor); down {
		return
	}

	w.log.Plain("poll hypervisor %s failed: %v", hypervisor, e)
}
