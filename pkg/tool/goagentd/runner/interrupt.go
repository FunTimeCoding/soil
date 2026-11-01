package runner

func (r *Runner) Interrupt() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.interrupted = true
	r.kill()
}
