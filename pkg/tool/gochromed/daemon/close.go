package daemon

func (d *Daemon) Close() {
	d.guard.Close()
	d.source.Close()
}
