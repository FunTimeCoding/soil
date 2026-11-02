package daemon

func (d *Daemon) Close() {
	d.pool.Close()

	if d.netbox == nil {
		return
	}

	d.netboxGuard.Close()
	d.netbox.Close()
}
