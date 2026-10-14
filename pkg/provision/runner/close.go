package runner

func (r *Runner) Close() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.closed {
		return
	}

	r.closed = true
	close(r.stop)

	if !r.started {
		close(r.finished)
	}
}
