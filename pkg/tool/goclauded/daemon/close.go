package daemon

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational"
)

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.stop.Close()

			for _, c := range d.clocks {
				c.Close()
			}

			d.gauge.Close()
			d.watcher.Close()
			d.memoryGuard.Close()
			d.queryGuard.Close()

			if d.browser != nil {
				d.browser.Close()
			}

			d.memory.Close()
			d.query.Close()
			errors.PanicClose(d.searchDatabase)
			relational.CloseMapper(d.database)
		},
	)
}
