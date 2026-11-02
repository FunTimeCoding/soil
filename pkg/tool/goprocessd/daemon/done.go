package daemon

func (d *Daemon) Done() <-chan struct{} {
	return d.supervisor.Done()
}
