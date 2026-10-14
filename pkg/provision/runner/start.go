package runner

func (r *Runner) Start() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.closed {
		return
	}

	r.started = true
	go r.run()
}
