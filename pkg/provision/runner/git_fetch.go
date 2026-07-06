package runner

func (r *Runner) gitFetch() {
	c := r.newRun()
	c.Directory = r.clonePath
	c.Start("git", "fetch", "origin")
}
