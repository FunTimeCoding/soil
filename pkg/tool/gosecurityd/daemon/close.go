package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			if d.feedClock != nil {
				d.feedClock.Close()
				d.ansibleGuard.Close()
				d.ansible.Close()
			}

			d.scoreClock.Close()
			d.clock.Close()
			d.extendedGuard.Close()
			d.trackerGuard.Close()
			d.nationalGuard.Close()
			d.extendedTracker.Close()
			d.securityTracker.Close()
			d.nationalDatabase.Close()
			relational.CloseMapper(d.database)
		},
	)
}
