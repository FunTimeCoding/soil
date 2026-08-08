package runner

func (r *Runner) gitConfigure() {
	c := r.newRun()
	c.Directory = r.clonePath
	c.Start("git", "config", "core.filemode", "false")
}
