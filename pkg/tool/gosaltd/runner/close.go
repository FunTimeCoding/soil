package runner

func (r *Runner) Close() {
	r.provision.Close()
}
