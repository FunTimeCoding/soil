package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			if d.clock != nil {
				d.clock.Close()
			}

			d.guard.Close()
			d.client.Close()
			relational.CloseMapper(d.database)
		},
	)
}
