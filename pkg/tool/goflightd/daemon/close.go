package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			for _, c := range d.clocks {
				c.Close()
			}

			relational.CloseMapper(d.database)
		},
	)
}
