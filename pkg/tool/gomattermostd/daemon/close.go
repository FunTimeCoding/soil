package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			for _, c := range d.clocks {
				c.Close()
			}

			if d.watcher != nil {
				d.watcher.Close()
				d.service.FlushDrained(d.watcher.Drain())
				d.rosterGuard.Close()
				d.connector.Close()
			}

			d.guard.Close()
			d.mattermost.Close()
			relational.CloseMapper(d.database)
		},
	)
}
