package runner

func (r *Runner) gitFetch() {
	c := r.newRun()
	c.Directory = r.cloneDirectory
	c.Start("git", "fetch", "origin")
}
