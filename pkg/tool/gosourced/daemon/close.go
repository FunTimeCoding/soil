package daemon

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.guard.Close()
			d.memory.Close()
		},
	)
}
