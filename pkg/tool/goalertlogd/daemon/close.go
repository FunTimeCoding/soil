package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.clock.Close()
			d.guard.Close()
			d.alertmanager.Close()
			relational.CloseMapper(d.database)
		},
	)
}
