package runner

func (r *Runner) validRepository() bool {
	c := r.newRun().NoPanic()
	c.Directory = r.clonePath
	c.Start("git", "rev-parse", "--is-inside-work-tree")

	return c.Error == nil
}
