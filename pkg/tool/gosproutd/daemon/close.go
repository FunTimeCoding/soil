package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.watcher.Close()
			relational.CloseMapper(d.database)
		},
	)
}
