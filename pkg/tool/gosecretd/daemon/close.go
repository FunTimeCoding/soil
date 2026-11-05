package daemon

func (d *Daemon) Close() {
	d.watch.Close()
	d.guard.Close()
	d.client.Close()
}
