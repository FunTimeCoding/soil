package runner

func (r *Runner) Close() {
	r.provision.Close()

	if r.inventory != nil {
		r.inventory.Close()
	}
}
