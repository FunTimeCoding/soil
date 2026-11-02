package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.clock.Close()
			d.netboxGuard.Close()

			for _, g := range d.guards {
				g.Close()
			}

			d.netbox.Close()

			for _, p := range d.peers {
				p.Close()
			}

			d.authorization.Close()
			relational.CloseMapper(d.database)
		},
	)
}
