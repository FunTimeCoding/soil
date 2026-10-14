package runner

func (r *Runner) Interrupt() {
	r.provision.Interrupt()
}
