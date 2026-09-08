package runner

func (r *Runner) healRepository() {
	if r.validRepository() {
		return
	}

	r.gitClone()
}
