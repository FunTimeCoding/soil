package daemon

func (d *Daemon) Close() {
	d.clock.Close()
	d.guard.Close()
	d.client.Close()
}
