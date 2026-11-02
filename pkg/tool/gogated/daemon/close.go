package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			if d.directoryGuard != nil {
				d.directoryGuard.Close()
			}

			d.authorization.Close()
			relational.CloseMapper(d.database)
		},
	)
}
