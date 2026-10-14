package runner

func (r *Runner) Interrupt() {
	if r.registry != nil {
		r.registry.Interrupt()
	}
}
