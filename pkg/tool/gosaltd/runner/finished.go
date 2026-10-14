package runner

func (r *Runner) Finished() <-chan struct{} {
	return r.provision.Finished()
}
