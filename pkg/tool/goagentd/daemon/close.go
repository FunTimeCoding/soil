package daemon

func (d *Daemon) Close() {
	d.runner.Close()
}
