package runner

func (r *Runner) Start() {
	r.provision.Start()
	r.keySyncing.Add(1)
	go r.keySyncLoop()
}
