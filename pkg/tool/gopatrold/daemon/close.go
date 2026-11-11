package daemon

func (d *Daemon) Close() {
	d.clock.Close()

	for _, g := range d.guards {
		g.Close()
	}
}
