package runner

func (r *Runner) Close() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.closed {
		return
	}

	r.closed = true
	close(r.stop)
	r.cancel()

	if !r.started {
		close(r.finished)
	}
}
