package worker

func (w *Worker) Unreachable(hypervisor string) bool {
	down, _ := w.reacher.Down(hypervisor)

	return down
}
