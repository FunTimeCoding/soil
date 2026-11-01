package runner

func (r *Runner) end() {
	r.ending.Do(func() { close(r.finished) })
}
