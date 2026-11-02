package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.runner.Close()
			<-d.runner.Finished()

			if d.inventory != nil {
				d.inventoryGuard.Close()
				d.inventory.Close()
			}

			d.reaper.Close()
			relational.CloseMapper(d.database)
		},
	)
}
