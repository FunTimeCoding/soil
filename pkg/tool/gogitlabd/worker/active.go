package worker

func (w *Worker) Active() bool {
	for _, entry := range w.Entries() {
		if entry.Active() {
			return true
		}
	}

	return false
}
