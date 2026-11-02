package daemon

func (d *Daemon) Close() {
	d.worker.Close()
	d.pool.Close()
}
