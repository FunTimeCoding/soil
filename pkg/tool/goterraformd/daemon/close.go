package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.runner.Close()
			<-d.runner.Finished()

			for _, g := range d.guards {
				g.Close()
			}

			for _, c := range d.downstream {
				c.Close()
			}

			d.reaper.Close()
			relational.CloseMapper(d.database)
		},
	)
}
