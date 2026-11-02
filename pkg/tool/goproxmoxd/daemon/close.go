package daemon

func (d *Daemon) Close() {
	d.worker.Close()

	if d.netbox != nil {
		d.reconciler.Close()
		d.netboxGuard.Close()
		d.netbox.Close()
	}

	d.shells.Close()
	d.pool.Close()
}
