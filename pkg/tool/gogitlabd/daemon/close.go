package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			for _, c := range d.clocks {
				c.Close()
			}

			d.worker.Close()

			if d.connector != nil {
				d.rosterGuard.Close()
				d.connector.Close()
			}

			d.pool.Close()
			relational.CloseMapper(d.database)
		},
	)
}
