package daemon

func (d *Daemon) Close() {
	d.guard.Close()
	d.client.Close()
}
