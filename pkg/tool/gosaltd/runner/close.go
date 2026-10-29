package runner

func (r *Runner) Close() {
	r.provision.Close()
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.closed = true

	if r.built != nil {
		r.built.Close()
	}
}
