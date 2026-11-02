package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.forwards.Close()

			for _, g := range d.guards {
				g.Close()
			}

			relational.CloseMapper(d.database)
		},
	)
}
