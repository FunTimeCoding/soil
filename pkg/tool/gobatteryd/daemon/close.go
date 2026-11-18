package daemon

func (d *Daemon) Close() {
	d.clock.Close()

	for _, t := range d.targets {
		t.Close()
	}
}
