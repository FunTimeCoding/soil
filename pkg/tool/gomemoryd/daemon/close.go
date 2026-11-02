package daemon

import "github.com/funtimecoding/soil/pkg/errors"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.clock.Close()
			d.guard.Close()
			d.query.Close()
			errors.PanicClose(d.database)
		},
	)
}
