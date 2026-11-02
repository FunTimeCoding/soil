package daemon

func (d *Daemon) Close() {
	d.closeOnce.Do(d.client.Close)
}
