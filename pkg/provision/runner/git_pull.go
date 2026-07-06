package runner

func (r *Runner) gitPull() {
	c := r.newRun()
	c.Directory = r.clonePath
	c.Start("git", "pull", "origin", Branch)
}
