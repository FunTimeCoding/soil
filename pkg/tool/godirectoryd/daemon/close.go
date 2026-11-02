package daemon

func (d *Daemon) Close() {
	d.guard.Close()
	d.authorization.Close()
}
