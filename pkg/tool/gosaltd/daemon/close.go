package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.runner.Close()
			<-d.runner.Finished()
			d.salt.Close()
			d.reaper.Close()
			relational.CloseMapper(d.database)
		},
	)
}
