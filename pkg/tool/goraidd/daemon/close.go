package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.watcher.Close()
			d.parserGuard.Close()
			d.parser.Close()
			d.authorization.Close()
			relational.CloseMapper(d.database)
		},
	)
}
