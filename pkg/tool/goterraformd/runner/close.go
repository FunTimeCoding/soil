package runner

func (r *Runner) Close() {
	r.provision.Close()

	for _, c := range r.downstream {
		c.Close()
	}
}
