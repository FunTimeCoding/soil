package daemon

func (d *Daemon) Close() {
	d.pool.Close()
}
