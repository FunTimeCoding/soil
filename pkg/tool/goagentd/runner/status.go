package runner

func (r *Runner) Status() *Result {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return &Result{State: r.state, Result: r.result}
}
