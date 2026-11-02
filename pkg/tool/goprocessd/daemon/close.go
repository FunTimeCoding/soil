package daemon

func (d *Daemon) Close() {
	d.supervisor.Close()
}
